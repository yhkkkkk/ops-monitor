package services

import (
	"fmt"
	"ops-monitor/pkg/ctx"
	"ops-monitor/pkg/logger"
	"ops-monitor/pkg/tools"
)

type (
	userPermissionService struct {
		ctx *ctx.Context
	}

	InterUserPermissionService interface {
		List() (interface{}, interface{})
	}
)

func newInterUserPermissionService(ctx *ctx.Context) InterUserPermissionService {
	return &userPermissionService{
		ctx: ctx,
	}
}

func (up userPermissionService) List() (interface{}, interface{}) {
	lc := logger.NewLogContext().
		WithTraceID(tools.GetTraceID(up.ctx.Ctx)).
		WithAction("获取权限列表")
	data, err := up.ctx.DB.UserPermissions().List()
	if err != nil {
		logger.Error(up.ctx.Ctx, lc, fmt.Errorf("查询权限列表失败: %v", err))
		return nil, err
	}
	logger.Info(up.ctx.Ctx, lc.WithParams(map[string]interface{}{
		"permission_count": len(data),
	}))

	return data, nil
}
