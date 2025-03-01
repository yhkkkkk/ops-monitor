package api

import (
	middleware "ops-monitor/internal/middleware"
	"ops-monitor/internal/models"
	"ops-monitor/internal/services"

	"github.com/gin-gonic/gin"
)

type AlertEventController struct{}

// API 告警事件相关路由
// @Summary 告警事件路由组
// @Description 告警事件相关的API路由组
// @Tags 告警事件
// @BasePath /api/ops/event
func (e AlertEventController) API(gin *gin.RouterGroup) {
	event := gin.Group("event")
	event.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
	)
	{
		event.GET("curEvent", e.ListCurrentEvent)
		event.GET("hisEvent", e.ListHistoryEvent)
	}
}

// ListCurrentEvent 获取当前告警事件列表
// @Summary 获取当前告警事件
// @Description 获取当前活跃的告警事件列表
// @Tags 告警事件
// @Accept json
// @Produce json
// @Param tenant_id header string true "租户ID"
// @Param query query models.ReqAlertCurEventQuery false "查询参数"
// @Success 200 {object} response.ResponseData{data=[]models.AlertCurEvent} "成功"
// @Failure 400 {object} response.ResponseData "请求参数错误"
// @Failure 401 {object} response.ResponseData "未授权"
// @Failure 500 {object} response.ResponseData "服务器内部错误"
// @Router /api/ops/event/curEvent [get]
func (e AlertEventController) ListCurrentEvent(ctx *gin.Context) {
	r := new(models.AlertCurEventQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.EventService.ListCurrentEvent(r)
	})
}

// ListHistoryEvent 获取历史告警事件列表
// @Summary 获取历史告警事件
// @Description 获取历史告警事件列表
// @Tags 告警事件
// @Accept json
// @Produce json
// @Param tenant_id header string true "租户ID"
// @Param query query models.ReqAlertCurEventQuery false "查询参数"
// @Success 200 {object} response.ResponseData{data=[]models.AlertHisEvent} "成功"
// @Failure 400 {object} response.ResponseData "请求参数错误"
// @Failure 401 {object} response.ResponseData "未授权"
// @Failure 500 {object} response.ResponseData "服务器内部错误"
// @Router /api/ops/event/hisEvent [get]
func (e AlertEventController) ListHistoryEvent(ctx *gin.Context) {
	r := new(models.AlertHisEventQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.EventService.ListHistoryEvent(r)
	})
}
