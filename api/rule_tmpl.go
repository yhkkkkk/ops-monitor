package api

import (
	middleware "ops-monitor/internal/middleware"
	"ops-monitor/internal/models"
	"ops-monitor/internal/services"

	"github.com/gin-gonic/gin"
)

type RuleTmplController struct{}

/*
规则模版 API
/api/ops/ruleTmpl
*/
func (rtc RuleTmplController) API(gin *gin.RouterGroup) {
	ruleTmplA := gin.Group("ruleTmpl")
	ruleTmplA.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
		middleware.AuditingLog(),
	)
	{
		ruleTmplA.POST("ruleTmplCreate", rtc.Create)
		ruleTmplA.POST("ruleTmplDelete", rtc.Delete)
	}

	ruleTmplB := gin.Group("ruleTmpl")
	ruleTmplB.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
	)
	{
		ruleTmplB.GET("ruleTmplList", rtc.List)
	}
}

func (rtc RuleTmplController) Create(ctx *gin.Context) {
	r := new(models.RuleTemplate)
	BindJson(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.RuleTmplService.Create(r)
	})
}

func (rtc RuleTmplController) Delete(ctx *gin.Context) {
	r := new(models.RuleTemplateQuery)
	BindJson(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.RuleTmplService.Delete(r)
	})
}

func (rtc RuleTmplController) List(ctx *gin.Context) {
	r := new(models.RuleTemplateQuery)
	BindQuery(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.RuleTmplService.List(r)
	})
}
