package routers

import (
	"asm_platform/infrastructure/pkg/slog"
	mds "asm_platform/interfaces/middelware"
	captcharouter "asm_platform/interfaces/routers/captcha"
	"asm_platform/interfaces/routers/login"
	oplog_router "asm_platform/interfaces/routers/oplog"
	user_router "asm_platform/interfaces/routers/user"
	"github.com/gin-gonic/gin"
)

// InitRouter 安装路由
func InitRouter(r *gin.Engine) *gin.Engine {
	r.MaxMultipartMemory = 120 * 1024 * 1024

	root := r.Group("/asm")
	root.Use(mds.GinLogger(slog.GetLogger()), mds.GinRecovery(slog.GetLogger(), true))
	{
		// Login
		login_router.SetupLoginRouter(root)

		// captcha
		captcharouter.SetupCaptchaRouter(root)

		// 授权
		root.Use(mds.JwtAuth())

		// logout refresh
		login_router.SetupLogOutRouter(root)

		// user
		user_router.SetupUserRouter(root)

		// 审计日志
		oplog_router.SetupOpLogRouter(root)

	}
	return r
}
