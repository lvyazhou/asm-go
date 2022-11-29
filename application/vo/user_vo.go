package vo

import "reflect"

// UserVO 用户返回VO
type UserVO struct {
	BasicVo

	// Account 账户名称
	Account string `json:"account"`

	// Name 姓名
	Name string `json:"name"`

	// Email 电子邮件
	Email string `json:"email"`

	// Mobile 手机号码
	Mobile string `json:"mobile"`

	// Status 用户状态 @see constants.user_status 1:正常；2：禁用；3：已删除'
	Status string `json:"status_name"`

	// 角色列表
	RoleList []*UserRoleVo `json:"role_list"`
}

func (user UserVO) IsEmpty() bool {
	return reflect.DeepEqual(user, UserVO{})
}
