package user_app

import (
	"asm_platform/application/app/role_app"
	user_dto "asm_platform/application/dto"
	user_vo "asm_platform/application/vo"
	user_entity "asm_platform/domain/entity/user"
	"asm_platform/infrastructure/pkg/constants"
	"asm_platform/infrastructure/pkg/slog"

	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	dao "asm_platform/infrastructure/pkg/database/mysql"
	"asm_platform/infrastructure/pkg/tool/utils"
	"asm_platform/infrastructure/repo"
	"github.com/gin-gonic/gin"
)

type UserApp struct {
}

func NewUserApp() UserAppInterface {
	return &UserApp{}
}

// userApp implements UserAppInterface
var _ UserAppInterface = &UserApp{}

// GetUserByUname
//  获取用户信息
//  @Description: 根据用户名/邮箱 查找用户
//  @params uName 用户名称
//  @return entity.User 用户实体信息
//  @return constapicode.SocError
//
func (u *UserApp) GetUserByUname(uName string) (*user_entity.User, constapicode.SocError) {
	d := repo.NewUserRepositoryDB()
	user, err := d.GetUserByAccount(&user_entity.User{
		Account: uName,
		Email:   uName,
	})
	if err != nil {
		slog.Errorf("get user by uname[%v] is error %v", uName, err)
		return nil, constapicode.ErrorOnDataFind
	}
	return user, constapicode.Success
}

// GetUser
//  获取用户信息
//  @Description: 根据用户ID查找用户
//  @params uid 用户ID
//  @return entity.User 用户实体信息
//  @return constapicode.SocError
//
func (u *UserApp) GetUser(uid int64) (*user_vo.UserVO, constapicode.SocError) {
	d := repo.NewUserRepositoryDB()
	user, err := d.GetUser(uid)
	if err != nil {
		return nil, constapicode.UserNotExist
	}
	userVo := user.UserEntityToVo()
	return userVo, constapicode.Success
}

// SaveUser
//  保存用户信息
//  @Description: 保存用户信息
//  @params user_dto.UserDTO 用户DTO
//  @return constapicode.SocError
//
func (u *UserApp) SaveUser(user *user_dto.UserDTO, c *gin.Context) constapicode.SocError {
	// 开启事务
	ctx := dao.Begin(c.Request.Context())
	d := repo.NewUserRepositoryDBT(ctx)
	// 保存用户
	id := utils_tool.GenerateUniqueId()
	userId := c.MustGet(constants.UserId).(int64)
	_, err := d.SaveUser(user.UserDtoToEntity(id, userId))
	// 保存用户和角色关系

	if err != nil {
		dao.Rollback(ctx)
	}
	// 事务提交
	dao.Commit(ctx)
	return constapicode.Success
}

// EditUser
//  编辑用户信息
//  @Description: 保存用户信息
//  @params user_dto.UserDTO 用户DTO
//  @return constapicode.SocError
//
func (u *UserApp) EditUser(user *user_dto.UserDTO, c *gin.Context) constapicode.SocError {
	//TODO implement me
	panic("implement me")
}

// EditUserStatus
//  编辑用户状态
//  @Description: 编辑用户状态
//  @params user_dto.UserStatusDto 用户状态DTO
//  @return constapicode.SocError
//
func (u *UserApp) EditUserStatus(user *user_dto.UserStatusDto, c *gin.Context) constapicode.SocError {
	//TODO implement me
	panic("implement me")
}

// DeleteUser
//  删除用户信息
//  @Description: 根据用户ID删除用户信息
//  @params uid 用户ID
//  @return constapicode.SocError
//
func (u *UserApp) DeleteUser(uid int64) constapicode.SocError {
	//TODO implement me
	panic("implement me")
}

// FindUserList
//  用户分页查询信息
//  @Description: 用户分页查询信息
//  @params user_dto.UserQueryDto 用户分页DTO
//  @return vo.UserVO 返回用户实体
//  @return constapicode.SocError
//

func (u *UserApp) FindUserList(userQueryDto *user_dto.UserQueryDto) ([]*user_vo.UserVO, int64, constapicode.SocError) {
	// 角色服务类
	roleApp := role_app.NewRoleApp()

	// 组装user query
	userQuery := userQueryDto.UserQueryDtoToEntity()

	// 用户查询
	d := repo.NewUserRepositoryDB()
	userList, totalCount, code := d.FindUserList(userQuery)
	if code != nil || len(userList) == 0 {
		return nil, 0, constapicode.ErrorOnDataFind
	}
	var userVoList []*user_vo.UserVO
	for u := range userList {
		user := userList[u]
		// convert to vo
		userVo := user.UserEntityToVo()
		// 查询用户角色
		userVo.RoleList = roleApp.FindUserRoleForUserId(user.ID)
		// 查询用户创建人员
		cuser, cerr := d.GetUser(user.CreateUser)
		if !cuser.IsEmpty() && cerr == nil {
			userVo.CreateUser = cuser.Name
		}
		uuser, uerr := d.GetUser(user.UpdateUser)
		if !uuser.IsEmpty() && uerr == nil {
			userVo.UpdateUser = uuser.Name
		}

		userVoList = append(userVoList, userVo)
	}

	return userVoList, totalCount, constapicode.Success
}

// EditUserPwd
//  用户密码修改
//  @Description: 修改用户密码
//  @params dto.UserPwdDto 用户密码DTO
//  @return constapicode.SocError
//
func (u *UserApp) EditUserPwd(userPwd *user_dto.UserPwdDto) constapicode.SocError {
	//TODO implement me
	panic("implement me")
}

// GetUserByMenuTree
//  用户权限查询
//  @Description: 个人用户权限查询，用于登录菜单授权；
//  @params vo.MenuNode 菜单权限vo
//  @return constapicode.SocError
//
func (u *UserApp) GetUserByMenuTree() (*user_vo.MenuNode, constapicode.SocError) {
	//TODO implement me
	panic("implement me")
}
