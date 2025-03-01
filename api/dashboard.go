package api

import (
	middleware "ops-monitor/internal/middleware"
	"ops-monitor/internal/models"
	"ops-monitor/internal/services"

	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

// API 仪表盘相关 API 路由
// @Summary 仪表盘路由组
// @Description 仪表盘相关的所有接口
// @Tags Dashboard
// @BasePath /api/ops/dashboard
func (dc DashboardController) API(gin *gin.RouterGroup) {
	dashboardA := gin.Group("dashboard")
	dashboardA.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
		middleware.AuditingLog(),
	)
	{
		dashboardA.POST("createDashboard", dc.Create)
		dashboardA.POST("updateDashboard", dc.Update)
		dashboardA.POST("deleteDashboard", dc.Delete)
		dashboardA.POST("createFolder", dc.CreateFolder)
		dashboardA.POST("updateFolder", dc.UpdateFolder)
		dashboardA.POST("deleteFolder", dc.DeleteFolder)
	}
	dashboardB := gin.Group("dashboard")
	dashboardB.Use(
		middleware.Auth(),
		middleware.Permission(),
		middleware.ParseTenant(),
	)
	{
		dashboardB.GET("getDashboard", dc.Get)
		dashboardB.GET("listDashboard", dc.List)
		dashboardB.GET("searchDashboard", dc.Search)
		dashboardB.GET("listFolder", dc.ListFolder)
		dashboardB.GET("getFolder", dc.GetFolder)
		dashboardB.GET("listGrafanaDashboards", dc.ListGrafanaDashboards)
		dashboardB.GET("getDashboardFullUrl", dc.GetDashboardFullUrl)
	}
}

// List 获取仪表盘列表
// @Summary 获取仪表盘列表
// @Description 获取指定租户下的所有仪表盘列表
// @Tags Dashboard
// @Accept json
// @Produce json
// @Param tenant_id query string true "租户ID"
// @Param query query models.DashboardQueryRequest false "查询参数"
// @Success 200 {object} response.ResponseData{data=[]models.Dashboard}
// @Router /api/ops/dashboard/listDashboard [get]
func (dc DashboardController) List(ctx *gin.Context) {
	r := new(models.DashboardQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.List(r)
	})
}

// Get 获取仪表盘详情
// @Summary 获取仪表盘详情
// @Description 获取指定仪表盘的详细信息
// @Tags Dashboard
// @Accept json
// @Produce json
// @Param tenant_id query string true "租户ID"
// @Param dashboard_id query string true "仪表盘ID"
// @Success 200 {object} response.ResponseData{data=models.Dashboard}
// @Router /api/ops/dashboard/getDashboard [get]
func (dc DashboardController) Get(ctx *gin.Context) {
	r := new(models.DashboardQuery)
	BindQuery(ctx, r)

	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.Get(r)
	})
}

// Create 创建仪表盘
// @Summary 创建仪表盘
// @Description 在指定租户下创建新的仪表盘
// @Tags Dashboard
// @Accept json
// @Produce json
// @Param tenant_id header string true "租户ID"
// @Param dashboard body models.DashboardCreateRequest true "仪表盘信息"
// @Success 200 {object} response.ResponseData{data=models.Dashboard}
// @Router /api/ops/dashboard/createDashboard [post]
func (dc DashboardController) Create(ctx *gin.Context) {
	r := new(models.Dashboard)
	BindJson(ctx, r)
	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)
	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.Create(r)
	})
}

// Update 更新仪表盘
// @Summary 更新仪表盘
// @Description 更新指定仪表盘的信息
// @Tags Dashboard
// @Accept json
// @Produce json
// @Param tenant_id header string true "租户ID"
// @Param dashboard body models.DashboardUpdateRequest true "仪表盘信息"
// @Success 200 {object} response.ResponseData{data=models.Dashboard}
// @Router /api/ops/dashboard/updateDashboard [post]
func (dc DashboardController) Update(ctx *gin.Context) {
	r := new(models.Dashboard)
	BindJson(ctx, r)
	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)
	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.Update(r)
	})
}

// Delete 删除仪表盘
// @Summary 删除仪表盘
// @Description 删除指定的仪表盘
// @Tags Dashboard
// @Accept json
// @Produce json
// @Param tenant_id header string true "租户ID"
// @Param query body models.DashboardQueryRequest true "删除条件"
// @Success 200 {object} response.ResponseData
// @Router /api/ops/dashboard/deleteDashboard [post]
func (dc DashboardController) Delete(ctx *gin.Context) {
	r := new(models.DashboardQuery)
	BindJson(ctx, r)
	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)
	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.Delete(r)
	})
}

// Search 搜索仪表盘
// @Summary 搜索仪表盘
// @Description 根据条件搜索仪表盘
// @Tags Dashboard
// @Accept json
// @Produce json
// @Param tenant_id query string true "租户ID"
// @Param query query models.DashboardQueryRequest false "搜索条件"
// @Success 200 {object} response.ResponseData{data=[]models.Dashboard}
// @Router /api/ops/dashboard/searchDashboard [get]
func (dc DashboardController) Search(ctx *gin.Context) {
	r := new(models.DashboardQuery)
	BindQuery(ctx, r)
	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)
	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.Search(r)
	})
}

// ListFolder 获取文件夹列表
// @Summary 获取文件夹列表
// @Description 获取指定租户下的所有仪表盘文件夹
// @Tags Dashboard-Folder
// @Accept json
// @Produce json
// @Param tenant_id query string true "租户ID"
// @Success 200 {object} response.ResponseData{data=[]models.DashboardFolders}
// @Router /api/ops/dashboard/listFolder [get]
func (dc DashboardController) ListFolder(ctx *gin.Context) {
	r := new(models.DashboardFolders)
	BindQuery(ctx, r)
	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.ListFolder(r)
	})
}

// GetFolder 获取文件夹详情
// @Summary 获取文件夹详情
// @Description 获取指定文件夹的详细信息
// @Tags Dashboard-Folder
// @Accept json
// @Produce json
// @Param tenant_id query string true "租户ID"
// @Param folder_id query string true "文件夹ID"
// @Success 200 {object} response.ResponseData{data=models.DashboardFolders}
// @Router /api/ops/dashboard/getFolder [get]
func (dc DashboardController) GetFolder(ctx *gin.Context) {
	r := new(models.DashboardFolders)
	BindQuery(ctx, r)
	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.GetFolder(r)
	})
}

// CreateFolder 创建文件夹
// @Summary 创建文件夹
// @Description 在指定租户下创建新的仪表盘文件夹
// @Tags Dashboard-Folder
// @Accept json
// @Produce json
// @Param tenant_id header string true "租户ID"
// @Param folder body models.FolderQuery true "文件夹信息"
// @Success 200 {object} response.ResponseData{data=models.DashboardFolders}
// @Router /api/ops/dashboard/createFolder [post]
func (dc DashboardController) CreateFolder(ctx *gin.Context) {
	r := new(models.DashboardFolders)
	BindJson(ctx, r)
	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.CreateFolder(r)
	})
}

// UpdateFolder 更新文件夹
// @Summary 更新文件夹
// @Description 更新指定文件夹的信息
// @Tags Dashboard-Folder
// @Accept json
// @Produce json
// @Param tenant_id header string true "租户ID"
// @Param folder body models.FolderUpdateRequest true "文件夹信息"
// @Success 200 {object} response.ResponseData{data=models.DashboardFolders}
// @Router /api/ops/dashboard/updateFolder [post]
func (dc DashboardController) UpdateFolder(ctx *gin.Context) {
	r := new(models.DashboardFolders)
	BindJson(ctx, r)
	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.UpdateFolder(r)
	})
}

// DeleteFolder 删除文件夹
// @Summary 删除文件夹
// @Description 删除指定的仪表盘文件夹
// @Tags Dashboard-Folder
// @Accept json
// @Produce json
// @Param tenant_id header string true "租户ID"
// @Param folder body models.FolderUpdateRequest true "文件夹信息"
// @Success 200 {object} response.ResponseData
// @Router /api/ops/dashboard/deleteFolder [post]
func (dc DashboardController) DeleteFolder(ctx *gin.Context) {
	r := new(models.DashboardFolders)
	BindJson(ctx, r)
	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.DeleteFolder(r)
	})
}

// ListGrafanaDashboards 获取 Grafana 仪表盘列表
// @Summary 获取 Grafana 仪表盘列表
// @Description 获取指定租户下的所有 Grafana 仪表盘
// @Tags Dashboard-Grafana
// @Accept json
// @Produce json
// @Param tenant_id query string true "租户ID"
// @Success 200 {object} response.ResponseData{data=[]models.DashboardFolders}
// @Router /api/ops/dashboard/listGrafanaDashboards [get]
func (dc DashboardController) ListGrafanaDashboards(ctx *gin.Context) {
	r := new(models.DashboardFolders)
	BindQuery(ctx, r)
	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.ListGrafanaDashboards(r)
	})
}

// GetDashboardFullUrl 获取仪表盘完整 URL
// @Summary 获取仪表盘完整 URL
// @Description 获取指定仪表盘的完整访问 URL
// @Tags Dashboard
// @Accept json
// @Produce json
// @Param tenant_id query string true "租户ID"
// @Param dashboard_id query string true "仪表盘ID"
// @Success 200 {object} response.ResponseData{data=string}
// @Router /api/ops/dashboard/getDashboardFullUrl [get]
func (dc DashboardController) GetDashboardFullUrl(ctx *gin.Context) {
	r := new(models.DashboardFolders)
	BindQuery(ctx, r)
	tid, _ := ctx.Get("TenantID")
	r.TenantId = tid.(string)

	Service(ctx, func() (interface{}, interface{}) {
		return services.DashboardService.GetDashboardFullUrl(r)
	})
}
