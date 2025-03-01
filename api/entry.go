package api

import (
	"ops-monitor/pkg/logger"
	"ops-monitor/pkg/response"
	"ops-monitor/pkg/tools"

	"github.com/gin-gonic/gin"
)

type ApiGroup struct {
	NoticeController
	DutyController
	DutyCalendarController
	CallbackController
	DatasourceController
	SilenceController
	RuleController
	UserController
	AlertEventController
	UserRoleController
	UserPermissionsController
	NoticeTemplateController
	RuleGroupController
	RuleTmplGroupController
	RuleTmplController
	DashboardInfoController
	TenantController
	DashboardController
	AuditLogController
	ClientController
	SettingsController
	SubscribeController
	ProbingController
}

var ApiGroupApp = new(ApiGroup)

func BindJson(ctx *gin.Context, req interface{}) {
	lc := logger.NewLogContext().
		WithTraceID(tools.GetTraceID(ctx.Request.Context())).
		WithAction("绑定JSON参数")

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx.Request.Context(), lc.WithParams(map[string]interface{}{
			"error": err.Error(),
		}), err)
		response.Fail(ctx, err.Error(), "failed")
		ctx.Abort()
		return
	}
}

func BindQuery(ctx *gin.Context, req interface{}) {
	lc := logger.NewLogContext().
		WithTraceID(tools.GetTraceID(ctx.Request.Context())).
		WithAction("绑定Query参数")

	err := ctx.ShouldBindQuery(req)
	if err != nil {
		logger.Error(ctx.Request.Context(), lc.WithParams(map[string]interface{}{
			"error": err.Error(),
		}), err)
		response.Fail(ctx, err.Error(), "failed")
		ctx.Abort()
		return
	}
}

func Service(ctx *gin.Context, fu func() (interface{}, interface{})) {
	data, err := fu()
	if err != nil {
		response.Fail(ctx, err.(error).Error(), "failed")
		ctx.Abort()
		return
	}
	response.Success(ctx, data, "success")
}

func ServiceWithLog(ctx *gin.Context, fu func() (interface{}, interface{})) {
	lc := logger.NewLogContext().
		WithTraceID(tools.GetTraceID(ctx.Request.Context()))

	data, err := fu()
	if err != nil {
		logger.Error(ctx.Request.Context(), lc.WithParams(map[string]interface{}{
			"error": err.(error).Error(),
		}), err.(error))
		response.Fail(ctx, err.(error).Error(), "failed")
		ctx.Abort()
		return
	}

	logger.Info(ctx.Request.Context(), lc.WithParams(map[string]interface{}{
		"response": "success",
	}))
	response.Success(ctx, data, "success")
}
