package api

import (
	middleware "ops-monitor/internal/middleware"
	"ops-monitor/internal/models"
	"ops-monitor/internal/services"
	jwtUtils "ops-monitor/pkg/tools"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

/*
用户 API
/api/ops/user
*/
func (uc UserController) API(gin *gin.RouterGroup) {
	userA := gin.Group("user")
	//userA.Use(
	//	middleware.Auth(),
	//	middleware.Permission(),
	//	middleware.ParseTenant(),
	//	middleware.AuditingLog(),
	//)
	{
		userA.POST("userUpdate", uc.Update)
		userA.POST("userDelete", uc.Delete)
		userA.POST("userChangePass", uc.ChangePass)
	}

	userB := gin.Group("user")
	userB.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
	)
	{
		userB.GET("userList", uc.List)
		userB.GET("searchDutyUser", uc.Search)
		userB.GET("searchUser", uc.Search)
	}
}

// List @Summary 更新用户信息
// @Description 获取用户信息接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseData "成功"
// @Failure 400 {object} response.ResponseData "失败"
// @Router /api/user/userList [get]
func (uc UserController) List(ctx *gin.Context) {
	r := new(models.MemberQuery)
	BindQuery(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserService.List(r)
	})
}

// Search @Summary 更新用户信息
// @Description 获取用户信息接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param username query string false "用户名"
// @Param email query string false "邮箱"
// @Param phone query string false "手机号"
// @Success 200 {object} response.ResponseData "更新成功"
// @Failure 400 {object} response.ResponseData "更新失败"
// @Router /api/user/searchUser [get]
// @Router /api/user/searchDutyUser [get]
func (uc UserController) Search(ctx *gin.Context) {
	r := new(models.MemberQuery)
	BindQuery(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserService.Search(r)
	})
}

// Get @Summary 更新用户信息
// @Description 获取用户信息接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param username query string false "用户名"
// @Param email query string false "邮箱"
// @Param phone query string false "手机号"
// @Success 200 {object} response.ResponseData "成功"
// @Failure 400 {object} response.ResponseData "失败"
// @Router /api/system/userInfo [get]
func (uc UserController) Get(ctx *gin.Context) {
	r := new(models.MemberQuery)
	token := ctx.Request.Header.Get("Authorization")
	username := jwtUtils.GetUser(token)
	r.UserName = username

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserService.Get(r)
	})
}

// Login @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "登录请求参数"
// @Success 200 {object} response.ResponseData "登录成功"
// @Failure 400 {object} response.ResponseData "请求失败"
// @Router /api/system/login [post]
func (uc UserController) Login(ctx *gin.Context) {
	r := new(models.Member)
	BindJson(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserService.Login(r)
	})
}

// Register @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body models.RegisterRequest true "登录请求参数"
// @Success 200 {object} response.ResponseData "登录成功"
// @Failure 400 {object} response.ResponseData "请求失败"
// @Router /api/system/register [post]
func (uc UserController) Register(ctx *gin.Context) {
	r := new(models.Member)
	BindJson(ctx, r)

	createUser := jwtUtils.GetUser(ctx.Request.Header.Get("Authorization"))
	r.CreateBy = createUser

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserService.Register(r)
	})
}

// Update @Summary 更新用户信息
// @Description 更新用户信息接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body models.UpdateRequest true "更新用户请求参数"
// @Success 200 {object} response.ResponseData "更新成功"
// @Failure 400 {object} response.ResponseData "更新失败"
// @Router /api/user/userUpdate [post]
func (uc UserController) Update(ctx *gin.Context) {
	r := new(models.Member)
	BindJson(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserService.Update(r)
	})
}

// Delete @Summary 更新用户信息
// @Description 更新用户信息接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body models.UpdateRequest true "更新用户请求参数"
// @Success 200 {object} response.ResponseData "更新成功"
// @Failure 400 {object} response.ResponseData "更新失败"
// @Router /api/user/userDelete [post]
func (uc UserController) Delete(ctx *gin.Context) {
	r := new(models.MemberQuery)
	BindJson(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserService.Delete(r)
	})
}

// CheckUser @Summary 更新用户信息
// @Description 更新用户信息接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param username query string false "用户名"
// @Param email query string false "邮箱"
// @Param phone query string false "手机号"
// @Success 200 {object} response.ResponseData "响应成功"
// @Failure 400 {object} response.ResponseData "响应失败"
// @Router /api/system/checkUser [get]
func (uc UserController) CheckUser(ctx *gin.Context) {
	r := new(models.MemberQuery)
	BindQuery(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserService.Get(r)
	})
}

// ChangePass @Summary 更新用户信息
// @Description 更新用户信息接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body models.ChangePassRequest true "更新用户请求参数"
// @Success 200 {object} response.ResponseData "更新成功"
// @Failure 400 {object} response.ResponseData "更新失败"
// @Router /api/user/changePass [post]
func (uc UserController) ChangePass(ctx *gin.Context) {
	r := new(models.Member)
	BindJson(ctx, r)

	Service(ctx, func() (interface{}, interface{}) {
		return services.UserService.ChangePass(r)
	})
}
