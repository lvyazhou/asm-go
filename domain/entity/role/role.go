package role_entity

import (
	"asm_platform/application/vo"
	"asm_platform/domain/entity"
	"reflect"
)

// Role 角色实体
type Role struct {
	entity.BasicEntity

	// Name 姓名
	Name string `gorm:"column:r_name"`

	// PlatFormId 所属平台
	PlatFormId int64 `gorm:"column:r_platform_id"`

	// Status 角色状态：1：启用；2：禁用；3：已删除
	Status int8 `gorm:"column:u_status"`
}

func (Role) TableName() string {
	return "u_role"
}

// IsEmpty 判断是否为空
func (role *Role) IsEmpty() bool {
	if role == nil {
		return false
	}
	return reflect.DeepEqual(role, &Role{})
}

// RoleEntityToUserRoleVo entity to vo
func (role *Role) RoleEntityToUserRoleVo() *vo.UserRoleVo {
	return &vo.UserRoleVo{
		RoleId:   role.ID,
		RoleName: role.Name,
	}
}

// RoleEntityToRoleVo entity to vo
func (role *Role) RoleEntityToRoleVo() *vo.RoleVo {
	return &vo.RoleVo{}
}
