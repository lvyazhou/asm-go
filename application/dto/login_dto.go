package dto

// LoginDTO 用户登录DTO
type LoginDTO struct {
	// 账号
	Account string `json:"account"`
	// 密码
	Password string `json:"password"`
	// 验证码
	AuthCode string `json:"auth_code"`
}
