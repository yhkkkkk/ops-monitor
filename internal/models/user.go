package models

// LoginRequest 登录请求参数
type LoginRequest struct {
	UserName string `json:"username" binding:"required" example:"admin"`  // 用户名
	Password string `json:"password" binding:"required" example:"123456"` // 密码
}

// RegisterRequest 注册请求参数
type RegisterRequest struct {
	UserName string `json:"username" binding:"required" example:"newuser"` // 用户名
	Password string `json:"password" binding:"required" example:"123456"`  // 密码
	Email    string `json:"email,omitempty" example:"user@example.com"`    // 邮箱
	Phone    string `json:"phone,omitempty" example:"13800138000"`         // 手机号
	CreateBy string `json:"create_by,omitempty" example:"system"`          // 创建人
}

// UpdateRequest 更新用户请求参数
type UpdateRequest struct {
	UserId   string `json:"user_id" binding:"required" example:"12345"` // 用户ID
	UserName string `json:"username,omitempty" example:"newname"`       // 用户名
	Email    string `json:"email,omitempty" example:"new@example.com"`  // 邮箱
	Phone    string `json:"phone,omitempty" example:"13800138000"`      // 手机号
}

// ChangePassRequest 修改密码请求参数
type ChangePassRequest struct {
	UserId      string `json:"user_id" binding:"required" example:"12345"`       // 用户ID
	NewPassword string `json:"new_password" binding:"required" example:"654321"` // 新密码
}

// UserQueryRequest 用户查询请求参数
type UserQueryRequest struct {
	UserName string `json:"username,omitempty" example:"admin"`   // 用户名
	Email    string `json:"email,omitempty" example:"@gmail.com"` // 邮箱
	Phone    string `json:"phone,omitempty" example:"138"`        // 手机号
}

type Member struct {
	UserId     string   `json:"userid"`
	UserName   string   `json:"username"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Password   string   `json:"password"`
	Role       string   `json:"role"`
	CreateBy   string   `json:"create_by"`
	CreateAt   int64    `json:"create_at"`
	JoinDuty   string   `json:"joinDuty" `
	DutyUserId string   `json:"dutyUserId"`
	Tenants    []string `json:"tenants" gorm:"tenants;serializer:json"`
}

type MemberQuery struct {
	UserId   string `json:"userid" form:"userid"`
	UserName string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Query    string `json:"query" form:"query"`
	JoinDuty string `json:"joinDuty" form:"joinDuty"`
}
