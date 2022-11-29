package user_app

import (
	user_dto "asm_platform/application/dto"
	user_vo "asm_platform/application/vo"
	user_entity "asm_platform/domain/entity/user"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	"github.com/gin-gonic/gin"
)

type UserAppInterface interface {
	// SaveUser 用户保存
	SaveUser(user *user_dto.UserDTO, c *gin.Context) constapicode.SocError

	// EditUser 用户编辑
	EditUser(user *user_dto.UserDTO, c *gin.Context) constapicode.SocError

	// EditUserStatus 用户状态更新
	EditUserStatus(user *user_dto.UserStatusDto, c *gin.Context) constapicode.SocError

	// DeleteUser 删除用户
	DeleteUser(uid int64) constapicode.SocError

	// GetUser 获取用户基本信息
	GetUser(uid int64) (*user_vo.UserVO, constapicode.SocError)

	// GetUserByUname 根据用户名/邮箱 查找用户
	GetUserByUname(uName string) (*user_entity.User, constapicode.SocError)

	// FindUserList 用户分页查询
	FindUserList(userQuery *user_dto.UserQueryDto) ([]*user_vo.UserVO, int64, constapicode.SocError)

	// EditUserPwd 修改密码
	EditUserPwd(userPwd *user_dto.UserPwdDto) constapicode.SocError

	// GetUserByMenuTree 获取当前人员菜单权限
	GetUserByMenuTree() (*user_vo.MenuNode, constapicode.SocError)
}
