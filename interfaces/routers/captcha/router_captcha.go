package captcha_router

import (
	captchaapi "asm_platform/interfaces/handler/captcha"
	mds "asm_platform/interfaces/middelware"
	"github.com/gin-gonic/gin"
)

func SetupCaptchaRouter(group *gin.RouterGroup) {
	group.Use(mds.Session("asm_platform-super-api"))
	group.GET("/captcha", func(c *gin.Context) {
		captchaapi.GetCaptchaPng(c, 4)
	})
}
