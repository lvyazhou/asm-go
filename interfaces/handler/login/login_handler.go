package login_handle

import (
	"asm_platform/application/app/login_app"
	"asm_platform/application/dto"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	"asm_platform/infrastructure/pkg/slog"
	"asm_platform/interfaces/handler"
	captchaapi "asm_platform/interfaces/handler/captcha"
	"github.com/gin-gonic/gin"
)

// requestLogin 登录请求
type requestLogin struct {
	// 账户名
	Account string `json:"account" binding:"required,min=6,max=64"`

	// 密码
	Password string `json:"password" binding:"required,min=8,max=32"`

	// 验证码
	AuthCode string `json:"auth_code" binding:"required,len=4"`
}

// Login godoc
// @Summary      登录
// @Description  登录
// @Tags         登录
// @Accept       json
// @Produce      json
// @Param        requestLogin  body      requestLogin  true  "登录请求信息"
// @Success      200           {object}  vo.LoginVo
// @Router       /login [post]
func (l *Logins) Login(c *gin.Context) {
	var rd requestLogin
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		handler.ReturnFormat(c, constapicode.ErrorReq, nil)
		return
	}
	// 校验验证码
	if captchaapi.CaptchaVerify(c, rd.AuthCode) {
		slog.Infof("[login][request][/login/ [post]] %v valid success.", rd.AuthCode)
	} else {
		handler.ReturnFormat(c, constapicode.ErrorVerifyToken, nil)
		return
	}
	// 登录逻辑
	data, code := l.us.LoginUser(&dto.LoginDTO{
		Account:  rd.Account,
		Password: rd.Password,
		AuthCode: rd.AuthCode,
	}, c)
	handler.ReturnFormat(c, code, data)
	return
}

// LogOut godoc
// @Summary      登出
// @Description  登出
// @Tags         登出
// @Accept       json
// @Produce      json
// @Success      200  {object}  handler.Response
// @Router       /logout [post]
// @Security     ApiKeyAuth
func (l *Logins) LogOut(c *gin.Context) {
	code := l.us.LoginOut(c)
	handler.ReturnFormat(c, code, nil)
	return
}

// Refresh godoc
// @Summary      用户续期
// @Description  用户续期
// @Tags         用户续期
// @Accept       json
// @Produce      json
// @Success      200           {object}  vo.LoginVo
// @Router       /refresh [post]
// @Security     ApiKeyAuth
func (l *Logins) Refresh(c *gin.Context) {
	loginVo, code := l.us.Refresh(c)
	handler.ReturnFormat(c, code, loginVo)
	return
}

type Logins struct {
	us login_app.LoginAppInterface
}

func NewLogins(us login_app.LoginAppInterface) *Logins {
	return &Logins{us: us}
}
