package api

import (
	middleware "ops-monitor/internal/middleware"
	"ops-monitor/internal/models"
	"ops-monitor/internal/services"

	"github.com/gin-gonic/gin"
)

type AuditLogController struct{}

func (ac AuditLogController) API(gin *gin.RouterGroup) {
	auditLog := gin.Group("auditLog")
	auditLog.Use(
		middleware.Cors(),
		middleware.Auth(),
		middleware.ParseTenant(),
	)
	{
		auditLog.GET("listAuditLog", ac.List)
		auditLog.GET("searchAuditLog", ac.Search)
	}
}

func (ac AuditLogController) List(ctx *gin.Context) {
	r := new(models.AuditLogQuery)
	BindQuery(ctx, r)
	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)
	Service(ctx, func() (interface{}, interface{}) {
		return services.AuditLogService.List(r)
	})
}

func (ac AuditLogController) Search(ctx *gin.Context) {
	r := new(models.AuditLogQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.AuditLogService.Search(r)
	})
}
