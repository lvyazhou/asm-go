package login_router

import (
	"asm_platform/application/app/login_app"
	login_handle "asm_platform/interfaces/handler/login"
	mds "asm_platform/interfaces/middelware"
	"github.com/gin-gonic/gin"
)

func SetupLoginRouter(group *gin.RouterGroup) {
	group.Use(mds.Session("asm_platform-super-api"))
	// 实例化应用层接口
	login := login_handle.NewLogins(login_app.NewLoginApp())
	group.POST("/login", login.Login)
	group.POST("/refresh", login.Refresh)
}

func SetupLogOutRouter(group *gin.RouterGroup) {
	group.Use(mds.Session("asm_platform-super-api"))
	// 实例化应用层接口
	login := login_handle.NewLogins(login_app.NewLoginApp())
	group.POST("/logout", login.LogOut)
}
