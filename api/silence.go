package api

import (
	middleware "ops-monitor/internal/middleware"
	"ops-monitor/internal/models"
	"ops-monitor/internal/services"
	jwtUtils "ops-monitor/pkg/tools"

	"github.com/gin-gonic/gin"
)

type SilenceController struct{}

/*
告警静默 API
/api/ops/silence
*/
func (sc SilenceController) API(gin *gin.RouterGroup) {
	silenceA := gin.Group("silence")
	silenceA.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
		middleware.AuditingLog(),
	)
	{
		silenceA.POST("silenceCreate", sc.Create)
		silenceA.POST("silenceUpdate", sc.Update)
		silenceA.POST("silenceDelete", sc.Delete)
	}

	silenceB := gin.Group("silence")
	silenceB.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
	)
	{
		silenceB.GET("silenceList", sc.List)

	}
}

func (sc SilenceController) Create(ctx *gin.Context) {
	r := new(models.AlertSilences)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	user := jwtUtils.GetUser(ctx.Request.Header.Get("Authorization"))
	r.CreateBy = user

	Service(ctx, func() (interface{}, interface{}) {
		return services.SilenceService.Create(r)
	})
}

func (sc SilenceController) Update(ctx *gin.Context) {
	r := new(models.AlertSilences)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	user := jwtUtils.GetUser(ctx.Request.Header.Get("Authorization"))
	r.CreateBy = user

	Service(ctx, func() (interface{}, interface{}) {
		return services.SilenceService.Update(r)
	})
}

func (sc SilenceController) Delete(ctx *gin.Context) {
	r := new(models.AlertSilenceQuery)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.SilenceService.Delete(r)
	})
}

func (sc SilenceController) List(ctx *gin.Context) {
	r := new(models.AlertSilenceQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.SilenceService.List(r)
	})
}
