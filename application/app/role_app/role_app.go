package role_app

import (
	"asm_platform/application/vo"
	"asm_platform/infrastructure/repo"
)

type RoleApp struct {
}

func NewRoleApp() RoleAppInterface {
	return &RoleApp{}
}

// roleApp implements RoleAppInterface
var _ RoleAppInterface = &RoleApp{}

// FindUserRoleForUserId
// 获取角色信息
// @Description: 根据用户ID查找角色信息
// @params uid   用户ID
// @return vo.UserRoleVo 角色VO信息
//
func (r RoleApp) FindUserRoleForUserId(uid int64) []*vo.UserRoleVo {
	// 查询数据库信息
	d := repo.NewRoleRepositoryDB()
	roleList, err := d.FindRoleListByUserId(uid)
	// 定义返回结果
	var userRoleVoList []*vo.UserRoleVo

	if err == nil {
		for r := range roleList {
			role := roleList[r]
			// convert to vo
			userRoleVoList = append(userRoleVoList, role.RoleEntityToUserRoleVo())
		}
	}
	return userRoleVoList
}
