package role_app

import "asm_platform/application/vo"

type RoleAppInterface interface {
	// FindUserRoleForUserId  根据用户ID查询角色
	FindUserRoleForUserId(uid int64) []*vo.UserRoleVo
}
