package vo

// RoleVo 角色返回Vo
type RoleVo struct {
}

// UserRoleVo 用户角色vo
type UserRoleVo struct {
	RoleId   int64  `json:"role_id"`
	RoleName string `json:"role_name"`
}
