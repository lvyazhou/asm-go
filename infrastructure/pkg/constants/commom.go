// @Title  Title
// @Description:
// @Author: lvyazhou
// @Date: 2022/6/7 8:58

package constants

import mapset "github.com/deckarep/golang-set"

var (
	// header
	UserId     = "userId"
	UserName   = "userName"
	CustomId   = "customId"
	CustomName = "CustomName"

	// AsmOpLogIndex log es name
	AsmOpLogIndex = "asm_op_log"

	// CaptChaApiPath api filter
	CaptChaApiPath = "/asm/captcha"
	LoginApiPath   = "/asm/login"

	// FilterUserApiPath 过滤四种登录接口无法获取用户的情况
	FilterUserApiPath = mapset.NewSet(CaptChaApiPath, LoginApiPath)

	// FilterLoginApiPath 登录接口密码处理
	FilterLoginApiPath = mapset.NewSet(LoginApiPath)

	// FilterLogApiPath 过滤验证码及审计日志操作，不产生日志的情况；
	FilterLogApiPath = mapset.NewSet(CaptChaApiPath, "/asm/oplog/list/")

	// SessionID SESSIONID
	SessionID = "asmSessionId"
)
