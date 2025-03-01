package api

import (
	"ops-monitor/internal/middleware"
	"ops-monitor/internal/services"

	"github.com/gin-gonic/gin"
)

type UserPermissionsController struct{}

/*
用户权限 API
/api/ops/permissions
*/
func (urc UserPermissionsController) API(gin *gin.RouterGroup) {
	perms := gin.Group("permissions")
	perms.Use(
		middleware.Auth(),
	)
	{
		perms.GET("permsList", urc.List)
	}
}

// List @Summary 获取权限列表
// @Description 获取权限列表
// @Tags 权限管理
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseData "成功"
// @Failure 400 {object} response.ResponseData "失败"
// @Router /api/user/searchUser [get]
// @Router /api/user/searchDutyUser [get]
func (urc UserPermissionsController) List(ctx *gin.Context) {
	Service(ctx, func() (interface{}, interface{}) {
		return services.UserPermissionService.List()
	})
}
