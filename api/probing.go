package api

import (
	"ops-monitor/internal/middleware"
	"ops-monitor/internal/models"
	"ops-monitor/internal/services"

	"github.com/gin-gonic/gin"
)

type ProbingController struct{}

func (e ProbingController) API(gin *gin.RouterGroup) {
	eventA := gin.Group("probing")
	eventA.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
		middleware.AuditingLog(),
	)
	{
		eventA.POST("createProbing", e.Create)
		eventA.POST("updateProbing", e.Update)
		eventA.POST("deleteProbing", e.Delete)
		eventA.POST("onceProbing", e.Once)
	}

	eventB := gin.Group("probing")
	eventB.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
	)
	{
		eventA.GET("listProbing", e.List)
		eventA.GET("searchProbing", e.Search)
	}
}

func (e ProbingController) List(ctx *gin.Context) {
	r := new(models.ProbingRuleQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.ProbingService.List(r)
	})
}

func (e ProbingController) Search(ctx *gin.Context) {
	r := new(models.ProbingRuleQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.ProbingService.Search(r)
	})
}

func (e ProbingController) Create(ctx *gin.Context) {
	r := new(models.ProbingRule)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.ProbingService.Create(r)
	})
}

func (e ProbingController) Update(ctx *gin.Context) {
	r := new(models.ProbingRule)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.ProbingService.Update(r)
	})
}

func (e ProbingController) Delete(ctx *gin.Context) {
	r := new(models.ProbingRuleQuery)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.ProbingService.Delete(r)
	})
}

func (e ProbingController) Once(ctx *gin.Context) {
	r := new(models.OnceProbing)
	BindJson(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.ProbingService.Once(r)
	})
}
