package api

import (
	middleware "ops-monitor/internal/middleware"
	"ops-monitor/internal/models"
	"ops-monitor/internal/services"

	"github.com/gin-gonic/gin"
)

type RuleGroupController struct{}

/*
规则组 API
/api/ops/ruleGroup
*/
func (rc RuleGroupController) API(gin *gin.RouterGroup) {
	ruleGroupA := gin.Group("ruleGroup")
	ruleGroupA.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
		middleware.AuditingLog(),
	)
	{
		ruleGroupA.POST("ruleGroupCreate", rc.Create)
		ruleGroupA.POST("ruleGroupUpdate", rc.Update)
		ruleGroupA.POST("ruleGroupDelete", rc.Delete)
	}
	ruleGroupB := gin.Group("ruleGroup")
	ruleGroupB.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
	)
	{
		ruleGroupB.GET("ruleGroupList", rc.List)
	}
}

func (rc RuleGroupController) Create(ctx *gin.Context) {
	r := new(models.RuleGroups)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.RuleGroupService.Create(r)
	})
}

func (rc RuleGroupController) Update(ctx *gin.Context) {
	r := new(models.RuleGroups)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.RuleGroupService.Update(r)
	})
}

func (rc RuleGroupController) List(ctx *gin.Context) {
	r := new(models.RuleGroupQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.RuleGroupService.List(r)
	})
}

func (rc RuleGroupController) Delete(ctx *gin.Context) {
	r := new(models.RuleGroupQuery)
	BindJson(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.RuleGroupService.Delete(r)
	})
}
