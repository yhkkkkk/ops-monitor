package initialization

import (
	_ "ops-monitor/docs"
	"ops-monitor/internal/global"
	"ops-monitor/internal/middleware"
	"ops-monitor/internal/routers"
	v1 "ops-monitor/internal/routers/v1"

	// _ "net/http/pprof"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Ops Monitor API
// @version 1.0
// @description 运维监控系统 API 文档
// @host localhost:9001
// @BasePath /api
func InitRoute() {
	log.Info().Msg("服务启动")

	mode := global.Config.Server.Mode
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	ginEngine := gin.New()

	ginEngine.Use(
		gin.Recovery(),
		// 启用CORS中间件
		middleware.Cors(),
		// 启用链路追踪中间件
		middleware.TraceMiddleware(),
		// 自定义请求日志格式
		middleware.GinZapLogger(),
		//gin.LoggerWithFormatter(middleware.RequestLoggerFormatter),
	)
	if global.Config.Server.Mode != gin.ReleaseMode {
		ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// 注册pprof路由
	if global.Config.Server.EnablePprof {
		// ginEngine.GET("/debug/pprof/", gin.WrapH(http.DefaultServeMux))
		pprof.Register(ginEngine)
	}
	allRouter(ginEngine)

	err := ginEngine.Run(":" + global.Config.Server.Port)
	if err != nil {
		log.Error().Err(err).Msg("服务启动失败")
		return
	}
}

func allRouter(engine *gin.Engine) {
	routers.HealthCheck(engine)
	v1.Router(engine)
}
