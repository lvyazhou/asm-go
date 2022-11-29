package user_router

import (
	"asm_platform/application/app/user_app"
	user_handle "asm_platform/interfaces/handler/user"
	"github.com/gin-gonic/gin"
)

func SetupUserRouter(group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	{
		// 实例化应用层接口
		user := user_handle.NewUsers(user_app.NewUserApp())

		// save user
		userGroup.POST("/save/", user.SaveUser)

		// edit user
		userGroup.POST("/edit/", user.EditUser)

		// get user
		userGroup.GET("/detail/:id", user.GetUserById)

		// get user by token
		userGroup.GET("/get/", user.GetUser)

		// delete user id
		userGroup.DELETE("delete/:id", user.DeleteUser)

		// search user by page
		userGroup.POST("/list", user.ListUserByPage)

		// update user status
		userGroup.PUT("/status", user.EditUserStatus)

		// user id Tree menu list
		userGroup.GET("/tree", user.TreeMenuListByUserId)

		// user change pwd
		userGroup.PUT("/password", user.ChangePWD)
	}
}
