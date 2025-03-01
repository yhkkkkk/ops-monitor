package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// 设置全局日志格式
	zerolog.TimeFieldFormat = time.RFC3339Nano
	// 创建自定义的控制台写入器
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05.000",
		NoColor:    false,
		FormatFieldName: func(i interface{}) string {
			return fmt.Sprintf("%s=", i)
		},
		FormatFieldValue: func(i interface{}) string {
			switch v := i.(type) {
			case []byte:
				return string(v)
			case map[string]interface{}:
				b, _ := json.MarshalIndent(v, "", "  ")
				return string(b)
			case string:
				return v
			default:
				return fmt.Sprintf("%v", i)
			}
		},
	}
	log.Logger = log.Output(consoleWriter)
}

// LogContext 日志上下文结构
type LogContext struct {
	Module     string                 `json:"module"`      // 模块名
	Function   string                 `json:"function"`    // 函数名
	Action     string                 `json:"action"`      // 操作描述
	Params     map[string]interface{} `json:"params"`      // 请求参数
	RequestID  string                 `json:"request_id"`  // 请求ID
	UserID     string                 `json:"user_id"`     // 用户ID
	TraceID    string                 `json:"trace_id"`    // 链路追踪ID
	ClientIP   string                 `json:"client_ip"`   // 客户端IP
	UserAgent  string                 `json:"user_agent"`  // 用户代理
	StartTime  time.Time              `json:"start_time"`  // 开始时间
	Duration   time.Duration          `json:"duration"`    // 执行时长
	CallerInfo *CallerInfo            `json:"caller_info"` // 调用者信息
}

// CallerInfo 调用者信息
type CallerInfo struct {
	File       string `json:"file"`        // 文件名
	Line       int    `json:"line"`        // 行号
	Function   string `json:"function"`    // 函数名
	Package    string `json:"package"`     // 包名
	StackTrace string `json:"stack_trace"` // 堆栈信息
}

// getCallerInfo 获取调用者信息
func getCallerInfo(skip int) *CallerInfo {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return nil
	}

	// 获取函数信息
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return nil
	}

	// 获取包名和函数名
	fullName := fn.Name()
	packagePath := ""
	funcName := fullName
	if lastSlash := strings.LastIndex(fullName, "/"); lastSlash >= 0 {
		packagePath = fullName[:lastSlash]
	}
	if lastDot := strings.LastIndex(fullName, "."); lastDot >= 0 {
		funcName = fullName[lastDot+1:]
	}

	// 获取相对路径
	workDir, _ := filepath.Abs(".")
	relPath, _ := filepath.Rel(workDir, file)

	// 获取堆栈信息
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	stackTrace := string(buf[:n])

	return &CallerInfo{
		File:       relPath,
		Line:       line,
		Function:   funcName,
		Package:    packagePath,
		StackTrace: stackTrace,
	}
}

// Log 统一日志记录函数
func Log(ctx context.Context, level string, lc LogContext, err error) {
	// 获取调用者信息
	lc.CallerInfo = getCallerInfo(3)
	if !lc.StartTime.IsZero() {
		lc.Duration = time.Since(lc.StartTime)
	}

	// 创建日志事件
	var event *zerolog.Event
	switch strings.ToLower(level) {
	case "error":
		event = log.Error()
	case "info":
		event = log.Info()
	case "debug":
		event = log.Debug()
	case "warn":
		event = log.Warn()
	case "trace":
		event = log.Trace()
	case "fatal":
		event = log.Fatal()
	default:
		event = log.Info()
	}

	// 添加基础字段
	event.
		Str("module", lc.Module).
		Str("function", lc.Function).
		Str("action", lc.Action)

	// 添加错误信息
	if err != nil {
		event.Err(err)
	}

	// 添加调用者信息
	if lc.CallerInfo != nil {
		event.Str("caller", fmt.Sprintf("%s:%d", lc.CallerInfo.File, lc.CallerInfo.Line))
	}

	// 添加其他可选字段
	if len(lc.Params) > 0 {
		event.Interface("params", lc.Params)
	}
	if lc.RequestID != "" {
		event.Str("request_id", lc.RequestID)
	}
	if lc.UserID != "" {
		event.Str("user_id", lc.UserID)
	}
	if lc.Duration > 0 {
		event.Str("duration", lc.Duration.String())
	}
	if lc.ClientIP != "" {
		event.Str("client_ip", lc.ClientIP)
	}
	if lc.UserAgent != "" {
		event.Str("user_agent", lc.UserAgent)
	}

	// 添加链路追踪信息
	if lc.TraceID != "" {
		event.Str("trace_id", lc.TraceID)
	}

	// 发送日志
	event.Send()

	// 如果是错误级别且有堆栈信息，单独打印堆栈
	if strings.ToLower(level) == "error" && err != nil && lc.CallerInfo != nil && lc.CallerInfo.StackTrace != "" {
		log.Error().
			Str("type", "stack_trace").
			Err(err).
			Msg("\n" + lc.CallerInfo.StackTrace)
	}
}

// Error 记录错误日志
func Error(ctx context.Context, lc LogContext, err error) {
	Log(ctx, "error", lc, err)
}

// Info 记录信息日志
func Info(ctx context.Context, lc LogContext) {
	Log(ctx, "info", lc, nil)
}

// Debug 记录调试日志
func Debug(ctx context.Context, lc LogContext) {
	Log(ctx, "debug", lc, nil)
}

// Warn 记录调试日志
func Warn(ctx context.Context, lc LogContext) {
	Log(ctx, "warn", lc, nil)
}

// // NewLogContext 创建新的日志上下文
// func NewLogContext(module, function, action string) LogContext {
// 	return LogContext{
// 		Module:    module,
// 		Function:  function,
// 		Action:    action,
// 		Params:    make(map[string]interface{}),
// 		StartTime: time.Now(),
// 	}
// }

// NewLogContext 创建新的日志上下文，自动获取调用者信息
func NewLogContext() LogContext {
	// 获取调用者信息
	pc, _, _, ok := runtime.Caller(1)
	module := "Unknown"
	function := "Unknown"

	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			// 完整函数名，例如: ops-monitor/internal/services.userService.Get
			fullName := fn.Name()

			// 分割包名和函数名
			parts := strings.Split(fullName, ".")
			if len(parts) >= 2 {
				// 获取结构体名（如果存在）和方法名
				methodParts := parts[len(parts)-2:]
				function = methodParts[len(methodParts)-1]

				// 获取包名作为模块名
				pkgPath := strings.Split(fullName, ".")[0]
				pkgParts := strings.Split(pkgPath, "/")
				if len(pkgParts) > 0 {
					module = pkgParts[len(pkgParts)-1]
				}

				// 如果是结构体方法，添加结构体名
				if len(methodParts) > 1 {
					// 转换 userService 为 UserService
					structName := methodParts[0]
					if len(structName) > 0 {
						structName = strings.ToUpper(structName[:1]) + structName[1:]
					}
					module = structName
				}
			}
		}
	}

	return LogContext{
		Module:    module,
		Function:  function,
		Params:    make(map[string]interface{}),
		StartTime: time.Now(),
	}
}

// WithRequestID 链式调用方法
func (lc LogContext) WithRequestID(requestID string) LogContext {
	lc.RequestID = requestID
	return lc
}

func (lc LogContext) WithUserID(userID string) LogContext {
	lc.UserID = userID
	return lc
}

func (lc LogContext) WithClientInfo(clientIP, userAgent string) LogContext {
	lc.ClientIP = clientIP
	lc.UserAgent = userAgent
	return lc
}

func (lc LogContext) WithParams(params map[string]interface{}) LogContext {
	lc.Params = params
	return lc
}

func (lc LogContext) WithAction(action string) LogContext {
	lc.Action = action
	return lc
}

// WithTraceID 添加追踪ID
func (lc LogContext) WithTraceID(traceID string) LogContext {
	lc.TraceID = traceID
	return lc
}

// // Debugf 基于 LogContext 的快捷日志方法
// func Debugf(module, function, format string, v ...interface{}) {
// 	lc := NewLogContext(module, function, fmt.Sprintf(format, v...))
// 	Debug(context.Background(), lc)
// }

// // Infof 基于 LogContext 的快捷日志方法
// func Infof(module, function, format string, v ...interface{}) {
// 	lc := NewLogContext(module, function, fmt.Sprintf(format, v...))
// 	Info(context.Background(), lc)
// }

// // Warnf 基于 LogContext 的快捷日志方法
// func Warnf(module, function, format string, v ...interface{}) {
// 	lc := NewLogContext(module, function, fmt.Sprintf(format, v...))
// 	Log(context.Background(), "warn", lc, nil)
// }

// // Errorf 基于 LogContext 的快捷日志方法
// func Errorf(module, function, format string, err error, v ...interface{}) {
// 	lc := NewLogContext(module, function, fmt.Sprintf(format, v...))
// 	Log(context.Background(), "error", lc, err)
// }
