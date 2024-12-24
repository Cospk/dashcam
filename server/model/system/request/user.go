package request

// AddUser 添加用户
type AddUser struct {
	Username    string `json:"username" binding:"required"`     // 用户名
	Password    string `json:"password" binding:"required"`     // 密码
	Phone       string `json:"phone" binding:"omitempty,phone"` // 手机号
	Email       string `json:"email" binding:"omitempty,email"` // 邮箱
	Active      bool   `json:"active"`                          // 是否活跃
	RoleModelID uint   `json:"roleId" binding:"required"`       // 角色ID
}
