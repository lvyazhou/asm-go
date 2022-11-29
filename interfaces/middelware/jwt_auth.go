package mds

import (
	"asm_platform/infrastructure/pkg/constants"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	login_key "asm_platform/infrastructure/pkg/constants/redis_key/login"
	cache "asm_platform/infrastructure/pkg/database/redis"
	"asm_platform/infrastructure/pkg/slog"
	"asm_platform/infrastructure/pkg/tool/jwt"
	"asm_platform/interfaces/handler"
	"fmt"
	"github.com/gin-gonic/gin"
)

// JwtAuth 认证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(login_key.AsmToken)
		if token == "" {
			handler.ReturnFormat(c, constapicode.NotAuth, nil)
			c.Abort()
			return
		}

		claims, code := jwt_tool.ParseToken(token)
		if code != constapicode.Success {
			handler.ReturnFormat(c, code, nil)
			c.Abort()
			return
		}

		if err := claims.Valid(); err != nil {
			handler.ReturnFormat(c, constapicode.NotAuth, nil)
			c.Abort()
			return
		}

		loginKey := fmt.Sprintf(login_key.AsmLogin, claims.UserId)
		// 检查redis中是否存在
		isExist := cache.Exists(loginKey)
		if !isExist {
			slog.Error("[check exist token] token not exist in redis,key -> : ", loginKey)
			handler.ReturnFormat(c, constapicode.PleaseRepeatLogin, nil)
			c.Abort()
			return
		}

		c.Set(constants.UserId, claims.UserId)
		c.Set(constants.UserName, claims.Username)
		c.Next()

	}
}
