package login_app

import (
	"asm_platform/application/vo"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	jwt_tool "asm_platform/infrastructure/pkg/tool/jwt"
	utils_tool "asm_platform/infrastructure/pkg/tool/utils"
)

//
//  获取用户ID
//  @Description: 通过当前用户token获取用户ID
//  @params token 用户token
//  @return jwt_tool.SocClaims jwt实体信息
//  @return constapicode.SocError
//
func getUserIdByJwtToken(token string) (*jwt_tool.SocClaims, constapicode.SocError) {
	if token == "" {
		return nil, constapicode.NotAuth
	}

	claims, code := jwt_tool.ParseToken(token)
	if code != constapicode.Success {
		return nil, code
	}

	if err := claims.Valid(); err != nil {
		return nil, constapicode.NotAuth
	}
	return claims, constapicode.Success
}

//
//  获取用户信息
//  @Description: 解析用户信息
//  @params userStr 用户组装信息
//  @return vo.RefreshUser 用户信息
//
func getUserByRefreshToken(userStr string) *vo.RefreshUserVo {
	inf, code := utils_tool.Json2Struct(userStr)
	if code == constapicode.Success {
		user := inf.(vo.RefreshUserVo)
		if !user.IsEmpty() {
			return &user
		}
	}
	return nil
}
