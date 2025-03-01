package models

type Dashboard struct {
	TenantId    string `json:"tenantId"`
	ID          string `json:"id" `
	Name        string `json:"name" gorm:"unique"`
	URL         string `json:"url"`
	FolderId    string `json:"folderId"`
	Description string `json:"description"`
}

type DashboardQuery struct {
	TenantId string `json:"tenantId" form:"tenantId"`
	ID       string `json:"id" form:"id"`
	Query    string `json:"query" form:"query"`
}

type DashboardFolders struct {
	TenantId            string `json:"tenantId" form:"tenantId"`
	ID                  string `json:"id" form:"id"`
	Name                string `json:"name"`
	Theme               string `json:"theme" form:"theme"`
	GrafanaHost         string `json:"grafanaHost" form:"grafanaHost"`
	GrafanaFolderId     int    `json:"grafanaFolderId" form:"grafanaFolderId"`
	GrafanaDashboardUid string `json:"grafanaDashboardUid" form:"grafanaDashboardUid" gorm:"-"`
}

type GrafanaDashboardInfo struct {
	Uid      string `json:"uid"`
	Title    string `json:"title"`
	FolderId int    `json:"folderId"`
}

type GrafanaDashboardMeta struct {
	Meta meta `json:"meta"`
}

type meta struct {
	Url string `json:"url"`
}

// DashboardCreateRequest 创建仪表盘请求参数
type DashboardCreateRequest struct {
	Title       string `json:"title" binding:"required" example:"CPU监控"`       // 仪表盘标题
	FolderId    string `json:"folder_id" binding:"required" example:"folder1"` // 文件夹ID
	Description string `json:"description" example:"用于监控CPU使用率"`               // 描述信息
	Config      string `json:"config" binding:"required" example:"{...}"`      // 仪表盘配置
}

// DashboardUpdateRequest 更新仪表盘请求参数
type DashboardUpdateRequest struct {
	DashboardId string `json:"dashboard_id" binding:"required" example:"dash1"` // 仪表盘ID
	Title       string `json:"title,omitempty" example:"内存监控"`                  // 仪表盘标题
	Description string `json:"description,omitempty" example:"监控内存使用情况"`        // 描述信息
	Config      string `json:"config,omitempty" example:"{...}"`                // 仪表盘配置
}

// DashboardQueryRequest 仪表盘查询参数
type DashboardQueryRequest struct {
	DashboardId string `json:"dashboard_id,omitempty" form:"dashboard_id" example:"dash1"` // 仪表盘ID
	Title       string `json:"title,omitempty" form:"title" example:"CPU"`                 // 标题
	FolderId    string `json:"folder_id,omitempty" form:"folder_id" example:"folder1"`     // 文件夹ID
	TenantId    string `json:"tenant_id,omitempty" form:"tenant_id" example:"tenant1"`     // 租户ID
}

// FolderCreateRequest 创建文件夹请求参数
type FolderCreateRequest struct {
	Title       string `json:"title" binding:"required" example:"系统监控"` // 文件夹标题
	Description string `json:"description" example:"用于存放系统监控相关的仪表盘"`    // 描述信息
}

// FolderUpdateRequest 更新文件夹请求参数
type FolderUpdateRequest struct {
	FolderId    string `json:"folder_id" binding:"required" example:"folder1"` // 文件夹ID
	Title       string `json:"title,omitempty" example:"应用监控"`                 // 文件夹标题
	Description string `json:"description,omitempty" example:"应用监控相关仪表盘"`      // 描述信息
}

// FolderQuery 文件夹查询参数
type FolderQuery struct {
	FolderId string `json:"folder_id,omitempty" form:"folder_id" example:"folder1"` // 文件夹ID
	Title    string `json:"title,omitempty" form:"title" example:"系统监控"`            // 标题
	TenantId string `json:"tenant_id,omitempty" form:"tenant_id" example:"tenant1"` // 租户ID
}
