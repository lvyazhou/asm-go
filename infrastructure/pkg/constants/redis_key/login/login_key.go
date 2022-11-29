package login_key

var (
	// AsmLogin 存储登录jwt
	AsmLogin = "ASM:LOGIN:%d"

	// LoginHour 登录8小时
	LoginHour = 8

	// RefreshHour 最大过期7天，168小时
	RefreshHour = 168

	// AsmRefresh 存储最大过期时间key
	AsmRefresh = "ASM:REFRESH:%s"

	// AsmToken token前缀
	AsmToken = "ASM-TOKEN"
)
