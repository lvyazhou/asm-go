package captcha_handle

import (
	captcha_tool "asm_platform/infrastructure/pkg/tool/captcha"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CaptchaVerify 验证码验证
func CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// GetCaptchaPng godoc
// @Summary      验证码
// @Description  验证码
// @Tags         登录
// @Accept       json
// @Produce      json
// @Router       /captcha [get]
func GetCaptchaPng(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 140, 40
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaId := captcha.NewLen(l)
	session := sessions.Default(c)
	session.Set("captcha", captchaId)
	_ = session.Save()
	_ = captcha_tool.Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}
