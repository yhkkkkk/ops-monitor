package routers

import (
	"fmt"
	"net/http"
	"ops-monitor/internal/global"
	"ops-monitor/pkg/client"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

// HealthStatus 健康状态响应结构
type HealthStatus struct {
	Status       string            `json:"status"`       // 服务状态：UP/DOWN/WARN
	Time         time.Time         `json:"time"`         // 当前时间
	Dependencies map[string]string `json:"dependencies"` // 依赖服务状态
	System       SystemInfo        `json:"system"`       // 系统信息
	Resources    ResourceUsage     `json:"resources"`    // 资源使用情况
	Application  ApplicationInfo   `json:"application"`  // 应用信息
}

// SystemInfo 系统信息
type SystemInfo struct {
	Hostname      string    `json:"hostname"`       // 主机名
	Platform      string    `json:"platform"`       // 操作系统平台
	OS            string    `json:"os"`             // 操作系统信息
	KernelArch    string    `json:"kernel_arch"`    // 内核架构
	KernelVersion string    `json:"kernel_version"` // 内核版本
	GoVersion     string    `json:"go_version"`     // Go 版本
	NumCPU        int       `json:"num_cpu"`        // CPU 核心数
	Uptime        uint64    `json:"uptime"`         // 系统运行时间
	BootTime      time.Time `json:"boot_time"`      // 系统启动时间
}

// ResourceUsage 资源使用情况
type ResourceUsage struct {
	CPUUsage     float64   `json:"cpu_usage"`    // CPU 使用率
	CPULoad      []float64 `json:"cpu_load"`     // CPU 负载（1,5,15分钟）
	MemoryUsage  float64   `json:"memory_usage"` // 内存使用率
	MemoryTotal  uint64    `json:"memory_total"` // 总内存
	MemoryFree   uint64    `json:"memory_free"`  // 空闲内存
	MemoryUsed   uint64    `json:"memory_used"`  // 已用内存
	SwapUsage    float64   `json:"swap_usage"`   // 交换分区使用率
	SwapTotal    uint64    `json:"swap_total"`   // 交换分区总量
	DiskUsage    float64   `json:"disk_usage"`   // 磁盘使用率
	DiskTotal    uint64    `json:"disk_total"`   // 总磁盘空间
	DiskFree     uint64    `json:"disk_free"`    // 空闲磁盘空间
	NumGoroutine int       `json:"goroutines"`   // Goroutine 数量
	NumThreads   int32     `json:"threads"`      // 线程数
	GCPause      uint64    `json:"gc_pause"`     // GC 暂停时间
	GCRuns       uint32    `json:"gc_runs"`      // GC 运行次数
}

// ApplicationInfo 应用信息
type ApplicationInfo struct {
	Version       string    `json:"version"`         // 应用版本
	StartTime     time.Time `json:"start_time"`      // 启动时间
	Environment   string    `json:"environment"`     // 运行环境
	PID           int32     `json:"pid"`             // 进程ID
	MemoryUsed    uint64    `json:"memory_used"`     // 应用使用的内存
	NumFD         int32     `json:"num_fd"`          // 文件描述符数量
	CPUPercent    float64   `json:"cpu_percent"`     // CPU 使用百分比
	LastGC        time.Time `json:"last_gc"`         // 最后一次 GC 时间
	NextGC        uint64    `json:"next_gc"`         // 下次 GC 阈值
	GCCPUFraction float64   `json:"gc_cpu_fraction"` // GC CPU 使用比例
}

var (
	startTime = time.Now()
	// 资源使用阈值
	thresholds = struct {
		CPUUsage     float64
		MemoryUsage  float64
		DiskUsage    float64
		GCPause      uint64
		NumGoroutine int
	}{
		CPUUsage:     80,
		MemoryUsage:  80,
		DiskUsage:    85,
		GCPause:      100, // 毫秒
		NumGoroutine: 10000,
	}
)

// HealthCheck @Summary 健康检查
// @Description 检查服务的健康状态
// @Tags 系统
// @Accept json
// @Produce json
// @Success 200 {object} HealthStatus "健康状态"
// @Failure 503 {object} HealthStatus "服务不可用"
// @Router /health [get]
func HealthCheck(gin *gin.Engine) {
	gin.GET("health", health)
}

func health(ctx *gin.Context) {
	status := HealthStatus{
		Status:       "UP",
		Time:         time.Now(),
		Dependencies: checkDependencies(),
	}

	// 获取系统信息
	status.System = getSystemInfo()

	// 获取资源使用情况
	status.Resources = getResourceUsage()

	// 获取应用信息
	status.Application = getApplicationInfo()

	// 检查整体状态
	if !isHealthy(status) {
		status.Status = "DOWN"
		ctx.IndentedJSON(http.StatusServiceUnavailable, status)
		return
	}

	ctx.IndentedJSON(http.StatusOK, status)
}

func getSystemInfo() SystemInfo {
	info, _ := host.Info()
	bootTime := time.Unix(int64(info.BootTime), 0)

	return SystemInfo{
		Hostname:      info.Hostname,
		Platform:      info.Platform,
		OS:            info.OS,
		KernelArch:    info.KernelArch,
		KernelVersion: info.KernelVersion,
		GoVersion:     runtime.Version(),
		NumCPU:        runtime.NumCPU(),
		Uptime:        info.Uptime,
		BootTime:      bootTime,
	}
}

func getResourceUsage() ResourceUsage {
	usage := ResourceUsage{
		NumGoroutine: runtime.NumGoroutine(),
	}

	// CPU 信息
	if cpuPercent, err := cpu.Percent(time.Second, false); err == nil && len(cpuPercent) > 0 {
		usage.CPUUsage = cpuPercent[0]
	}

	// CPU 负载
	if loadInfo, err := load.Avg(); err == nil {
		usage.CPULoad = []float64{loadInfo.Load1, loadInfo.Load5, loadInfo.Load15}
	}

	// 内存信息
	if memInfo, err := mem.VirtualMemory(); err == nil {
		usage.MemoryUsage = memInfo.UsedPercent
		usage.MemoryTotal = memInfo.Total
		usage.MemoryFree = memInfo.Free
		usage.MemoryUsed = memInfo.Used
	}

	// 交换分区信息
	if swapInfo, err := mem.SwapMemory(); err == nil {
		usage.SwapUsage = swapInfo.UsedPercent
		usage.SwapTotal = swapInfo.Total
	}

	// 磁盘信息
	if diskInfo, err := disk.Usage("/"); err == nil {
		usage.DiskUsage = diskInfo.UsedPercent
		usage.DiskTotal = diskInfo.Total
		usage.DiskFree = diskInfo.Free
	}

	// 获取当前进程信息
	if proc, err := process.NewProcess(int32(os.Getpid())); err == nil {
		if numThreads, err := proc.NumThreads(); err == nil {
			usage.NumThreads = numThreads
		}
	}

	// GC 统计
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	usage.GCPause = m.PauseNs[(m.NumGC+255)%256] / 1e6 // 转换为毫秒
	usage.GCRuns = m.NumGC

	return usage
}

func getApplicationInfo() ApplicationInfo {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	appInfo := ApplicationInfo{
		Version:       global.Version,
		StartTime:     startTime,
		Environment:   gin.Mode(),
		PID:           int32(os.Getpid()),
		MemoryUsed:    m.Alloc,
		LastGC:        time.Unix(0, int64(m.LastGC)),
		NextGC:        m.NextGC,
		GCCPUFraction: m.GCCPUFraction,
	}

	// 获取进程详细信息
	if proc, err := process.NewProcess(appInfo.PID); err == nil {
		if numFD, err := proc.NumFDs(); err == nil {
			appInfo.NumFD = numFD
		}
		if cpuPercent, err := proc.CPUPercent(); err == nil {
			appInfo.CPUPercent = cpuPercent
		}
	}

	return appInfo
}

func checkDependencies() map[string]string {
	dependencies := make(map[string]string)

	// 检查数据库
	if err := checkDB(); err != nil {
		dependencies["database"] = "不可用: " + err.Error()
	} else {
		dependencies["database"] = "正常"
	}

	// 检查 Redis
	if err := checkRedis(); err != nil {
		dependencies["redis"] = "不可用: " + err.Error()
	} else {
		dependencies["redis"] = "正常"
	}

	return dependencies
}

// 检查数据库连接
func checkDB() error {
	db := client.GetDB()
	if db == nil {
		return fmt.Errorf("数据库未初始化")
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

// 检查 Redis 连接
func checkRedis() error {
	if client.Redis == nil {
		return fmt.Errorf("redis未初始化")
	}
	return client.Redis.Ping().Err()
}

func isHealthy(status HealthStatus) bool {
	// 检查 CPU 使用率
	if status.Resources.CPUUsage > thresholds.CPUUsage {
		return false
	}

	// 检查内存使用率
	if status.Resources.MemoryUsage > thresholds.MemoryUsage {
		return false
	}

	// 检查磁盘使用率
	if status.Resources.DiskUsage > thresholds.DiskUsage {
		return false
	}

	// 检查 Goroutine 数量
	if status.Resources.NumGoroutine > thresholds.NumGoroutine {
		return false
	}

	// 检查 GC 暂停时间
	if status.Resources.GCPause > thresholds.GCPause {
		return false
	}

	// 检查依赖服务
	for _, status := range status.Dependencies {
		if status != "正常" {
			return false
		}
	}

	return true
}
