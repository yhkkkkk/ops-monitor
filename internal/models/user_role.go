package models

type UserRole struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Permissions []UserPermissions `json:"permissions" gorm:"permissions;serializer:json"`
	CreateAt    int64             `json:"create_at"`
}

type UserRoleQuery struct {
	ID          string `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

// ReqUserRole 用户角色模型
type ReqUserRole struct {
	RoleId      string   `json:"role_id" example:"role_123"`                   // 角色ID
	RoleName    string   `json:"role_name" binding:"required" example:"admin"` // 角色名称
	Description string   `json:"description" example:"管理员角色"`                  // 角色描述
	Permissions []string `json:"permissions" example:"['read','write']"`       // 权限列表
	CreateBy    string   `json:"create_by" example:"system"`                   // 创建人
	CreateAt    int64    `json:"create_at" example:"1646064000"`               // 创建时间
}

// ReqUserRoleQuery 用户角色查询条件
type ReqUserRoleQuery struct {
	RoleId   string `json:"role_id" form:"role_id" example:"role_123"`  // 角色ID
	RoleName string `json:"role_name" form:"role_name" example:"admin"` // 角色名称
	Query    string `json:"query" form:"query" example:"管理员"`           // 模糊查询关键字
}
