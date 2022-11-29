package repository

import role_entity "asm_platform/domain/entity/role"

type RoleRepository interface {

	// FindRoleListByUserId 通过用户ID查询角色ID
	FindRoleListByUserId(userId int64) (result []*role_entity.Role, error error)
}
