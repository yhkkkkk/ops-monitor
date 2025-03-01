package api

import (
	middleware "ops-monitor/internal/middleware"
	"ops-monitor/internal/models"
	"ops-monitor/internal/services"

	"github.com/gin-gonic/gin"
)

type UserRoleController struct{}

// API 用户角色路由组
// @Summary 用户角色相关接口
// @Description 用户角色的增删改查接口
// @Tags 用户角色管理
// @BasePath /api/ops/role
func (urc UserRoleController) API(gin *gin.RouterGroup) {
	roleA := gin.Group("role")
	roleA.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
		middleware.AuditingLog(),
	)
	{
		roleA.POST("roleCreate", urc.Create)
		roleA.POST("roleUpdate", urc.Update)
		roleA.POST("roleDelete", urc.Delete)
	}

	roleB := gin.Group("role")
	roleB.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
	)
	{
		roleB.GET("roleList", urc.List)
	}
}

// Create 创建用户角色
// @Summary 创建用户角色
// @Description 创建新的用户角色
// @Tags 用户角色管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param role body models.ReqUserRole true "用户角色信息"
// @Success 200 {object} response.ResponseData{data=models.UserRole} "成功"
// @Failure 400 {object} response.ResponseData "请求参数错误"
// @Failure 401 {object} response.ResponseData "未授权"
// @Failure 500 {object} response.ResponseData "服务器内部错误"
// @Router /api/ops/role/roleCreate [post]
func (urc UserRoleController) Create(ctx *gin.Context) {
	r := new(models.UserRole)
	BindJson(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserRoleService.Create(r)
	})
}

// Update 更新用户角色
// @Summary 更新用户角色
// @Description 更新现有用户角色信息
// @Tags 用户角色管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param role body models.ReqUserRole true "用户角色信息"
// @Success 200 {object} response.ResponseData{data=models.UserRole} "成功"
// @Failure 400 {object} response.ResponseData "请求参数错误"
// @Failure 401 {object} response.ResponseData "未授权"
// @Failure 500 {object} response.ResponseData "服务器内部错误"
// @Router /api/ops/role/roleUpdate [post]
func (urc UserRoleController) Update(ctx *gin.Context) {
	r := new(models.UserRole)
	BindJson(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserRoleService.Update(r)
	})
}

// Delete 删除用户角色
// @Summary 删除用户角色
// @Description 删除指定的用户角色
// @Tags 用户角色管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param query body models.ReqUserRoleQuery true "用户角色查询条件"
// @Success 200 {object} response.ResponseData "成功"
// @Failure 400 {object} response.ResponseData "请求参数错误"
// @Failure 401 {object} response.ResponseData "未授权"
// @Failure 500 {object} response.ResponseData "服务器内部错误"
// @Router /api/ops/role/roleDelete [post]
func (urc UserRoleController) Delete(ctx *gin.Context) {
	r := new(models.UserRoleQuery)
	BindJson(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserRoleService.Delete(r)
	})
}

// List 获取用户角色列表
// @Summary 获取用户角色列表
// @Description 根据查询条件获取用户角色列表
// @Tags 用户角色管理
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param query query models.ReqUserRoleQuery false "查询条件"
// @Success 200 {object} response.ResponseData{data=[]models.UserRole} "成功"
// @Failure 400 {object} response.ResponseData "请求参数错误"
// @Failure 401 {object} response.ResponseData "未授权"
// @Failure 500 {object} response.ResponseData "服务器内部错误"
// @Router /api/ops/role/roleList [get]
func (urc UserRoleController) List(ctx *gin.Context) {
	r := new(models.UserRoleQuery)
	BindQuery(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserRoleService.List(r)
	})
}
