package jwt_tool

import (
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	login_key "asm_platform/infrastructure/pkg/constants/redis_key/login"
	utils_tool "asm_platform/infrastructure/pkg/tool/utils"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

type SocClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`

	jwt.StandardClaims
}

var jwtSecret = []byte("6190905e00434dd89cd3ad46e489c8f3")

func Init(secret string) {
	if secret != "" {
		jwtSecret = []byte(secret)
	}
}

// GeneratorToken 生成Token
func GeneratorToken(userId int64, username string) (token string, code constapicode.SocError) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(login_key.LoginHour) * time.Hour)

	claims := SocClaims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			//Audience: "",
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			Id:        strconv.FormatInt(utils_tool.GenerateUniqueId(), 10),
			//IssuedAt:  0,
			// 指定token发行人
			Issuer:    "lvyazhou",
			NotBefore: 0,
			Subject:   "asm",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		fmt.Printf("signed jwt token failed: %s", err.Error())
		return "", constapicode.ErrorGenerateToken
	}

	return token, constapicode.Success
}

// ParseToken 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string) (claims *SocClaims, code constapicode.SocError) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &SocClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, constapicode.ErrorParserToken
	}

	if tokenClaims == nil {
		return nil, constapicode.ErrorParserToken
	}

	// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
	// 要传入指针，项目中结构体都是用指针传递，节省空间。
	if claims, ok := tokenClaims.Claims.(*SocClaims); ok && tokenClaims.Valid {
		return claims, constapicode.Success
	}

	return nil, constapicode.ErrorParserToken
}
