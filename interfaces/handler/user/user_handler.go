package user_handle

import (
	"asm_platform/application/app/user_app"
	user_dto "asm_platform/application/dto"
	"asm_platform/infrastructure/pkg/constants"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	"asm_platform/infrastructure/pkg/slog"
	"asm_platform/interfaces/handler"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Users struct {
	us user_app.UserAppInterface
}

func NewUsers(us user_app.UserAppInterface) *Users {
	return &Users{us: us}
}

// SaveUser 创建用户
// @Summary      创建用户
// @Description  创建用户
// @accept       json
// @produce      json
// @Tags         用户管理
// @Param        userDTO  body      dto.UserDTO  true  "request user"
// @Success      200        {object}  handler.Response
// @Router       /user/save/ [post]
// @Security     ApiKeyAuth
func (u *Users) SaveUser(c *gin.Context) {
	var param = &user_dto.UserDTO{}
	if err := c.ShouldBindJSON(param); err != nil {
		slog.Errorf("[user][request][/user/ [post]] createUser valid error %v.", err.Error())
		handler.ReturnFormat(c, constapicode.ErrorReq, nil)
		return
	}
	u.us.SaveUser(param, c)
	handler.ReturnFormat(c, constapicode.Success, nil)
	return
}

// GetUserById godoc
// @Summary      个人用户信息
// @Description  个人用户信息
// @Tags         用户管理
// @Accept                       json
// @Param        id  path      int  true  "用户ID"
// @Success      200  {object}  handler.Response
// @Router       /user/get/{id} [get]
// @Security     ApiKeyAuth
func (u *Users) GetUserById(c *gin.Context) {
	uId, a := c.Params.Get("id")
	if !a {
		handler.ReturnFormat(c, constapicode.ErrorReq, nil)
		return
	}
	pId, _ := strconv.ParseInt(uId, 10, 64)
	user, code := u.us.GetUser(pId)
	handler.ReturnFormat(c, code, user)
	return
}

func (u *Users) EditUser(c *gin.Context) {

}

// GetUser godoc
// @Summary      个人用户信息
// @Description  个人用户信息
// @Tags         用户管理
// @Accept                       json
// @Success      200  {object}  vo.UserVO
// @Router       /user/get/ [get]
// @Security     ApiKeyAuth
func (u *Users) GetUser(c *gin.Context) {
	// 判断当前用户ID是否存在
	var userId int64
	if _, exists := c.Get(constants.UserId); exists {
		userId = c.MustGet(constants.UserId).(int64)
	} else {
		slog.Error("/usr/get/ not auth request")
		handler.ReturnFormat(c, constapicode.NotAuth, nil)
		return
	}
	user, code := u.us.GetUser(userId)
	handler.ReturnFormat(c, code, user)
	return
}

// DeleteUser godoc
// @Summary      删除用户信息
// @Description  根据用户ID删除用户信息
// @Tags         用户管理
// @Accept                       json
// @Success      200  {object}  handler.Response
// @Param        id  path      int  true  "用户ID"
// @Router       /user/del/{id} [delete]
// @Security     ApiKeyAuth
func (u *Users) DeleteUser(c *gin.Context) {

}

// ListUserByPage godoc
// @Summary      用户分页查询
// @Description  用户分页查询
// @Tags         用户管理
// @Accept                             json
// @Success      200  {object}         vo.UserVO
// @Param        userQuery   body      dto.UserQueryDto  true  "用户查询实体"
// @Router       /user/list/ [post]
// @Security     ApiKeyAuth
func (u *Users) ListUserByPage(c *gin.Context) {
	var param = &user_dto.UserQueryDto{}
	if err := c.ShouldBindJSON(param); err != nil {
		slog.Errorf("[user][request][/user/list/ [post]] list user by page valid error %v.", err.Error())
		handler.ReturnFormat(c, constapicode.ErrorReq, nil)
		return
	}
	ulist, totalCount, code := u.us.FindUserList(param)
	handler.ReturnPageFormat(c, code, ulist, totalCount)
	return
}

func (u *Users) EditUserStatus(c *gin.Context) {

}

func (u *Users) TreeMenuListByUserId(context *gin.Context) {

}

func (u *Users) ChangePWD(context *gin.Context) {

}
