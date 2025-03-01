package api

import (
	middleware "ops-monitor/internal/middleware"
	"ops-monitor/internal/models"
	"ops-monitor/internal/services"
	"ops-monitor/pkg/tools"

	"github.com/gin-gonic/gin"
)

type RuleController struct{}

/*
告警规则 API
/api/ops/rule
*/
func (rc RuleController) API(gin *gin.RouterGroup) {
	ruleA := gin.Group("rule")
	ruleA.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
		middleware.AuditingLog(),
	)
	{
		ruleA.POST("ruleCreate", rc.Create)
		ruleA.POST("ruleUpdate", rc.Update)
		ruleA.POST("ruleDelete", rc.Delete)
	}
	ruleB := gin.Group("rule")
	ruleB.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
	)
	{
		ruleB.GET("ruleList", rc.List)
		ruleB.GET("ruleSearch", rc.Search)
	}
}

func (rc RuleController) Create(ctx *gin.Context) {
	r := new(models.AlertRule)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)
	r.RuleId = "a-" + tools.RandId()

	Service(ctx, func() (interface{}, interface{}) {
		return services.RuleService.Create(r)
	})
}

func (rc RuleController) Update(ctx *gin.Context) {
	r := new(models.AlertRule)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.RuleService.Update(r)
	})
}

func (rc RuleController) List(ctx *gin.Context) {
	r := new(models.AlertRuleQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.RuleService.List(r)
	})
}

func (rc RuleController) Delete(ctx *gin.Context) {
	r := new(models.AlertRuleQuery)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.RuleService.Delete(r)
	})
}

func (rc RuleController) Search(ctx *gin.Context) {
	r := new(models.AlertRuleQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.RuleService.Search(r)
	})
}
