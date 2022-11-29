package login_app

import (
	"asm_platform/application/app/user_app"
	user_dto "asm_platform/application/dto"
	"asm_platform/application/vo"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	login_key "asm_platform/infrastructure/pkg/constants/redis_key/login"
	dao "asm_platform/infrastructure/pkg/database/mysql"
	cache "asm_platform/infrastructure/pkg/database/redis"
	"asm_platform/infrastructure/pkg/slog"
	jwt_tool "asm_platform/infrastructure/pkg/tool/jwt"
	slat_tool "asm_platform/infrastructure/pkg/tool/slat"
	utils_tool "asm_platform/infrastructure/pkg/tool/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type LoginApp struct {
}

func NewLoginApp() LoginAppInterface {
	return &LoginApp{}
}

// LoginApp implements LoginAppInterface
var _ LoginAppInterface = &LoginApp{}

type LoginAppInterface interface {
	// LoginUser 用户登录
	LoginUser(user *user_dto.LoginDTO, c *gin.Context) (*vo.LoginVo, constapicode.SocError)
	// LoginOut 用户登出
	LoginOut(c *gin.Context) constapicode.SocError
	// Refresh 用户续期
	Refresh(c *gin.Context) (*vo.LoginVo, constapicode.SocError)
}

// LoginUser 登录
func (loginApp *LoginApp) LoginUser(loginDTO *user_dto.LoginDTO, c *gin.Context) (*vo.LoginVo, constapicode.SocError) {
	// 开启事务
	ctx := dao.Begin(c.Request.Context())
	// 用户服务类
	userApp := user_app.NewUserApp()

	// 查询用户密码是否正确
	userEntity, code := userApp.GetUserByUname(loginDTO.Account)
	if code != constapicode.Success {
		dao.Rollback(ctx)
		return nil, code
	}
	// 判断用户是否存在
	if userEntity == nil {
		return nil, constapicode.UserNotExist
	}
	// 判断密码是否正确
	if userEntity.Password != utils_tool.Md5Crypto(slat_tool.SlatUserPwd+loginDTO.Password) {
		return nil, constapicode.ErrorOnLogin
	}
	// 生成jwt token
	token, code := jwt_tool.GeneratorToken(userEntity.ID, userEntity.Name)
	if code != constapicode.Success {
		return nil, code
	} else {
		// redis 缓存一份
		loginKey := fmt.Sprintf(login_key.AsmLogin, userEntity.ID)
		cache.Set(loginKey, token, time.Duration(login_key.LoginHour)*time.Hour)

		// redis 缓存token，存储userId，目的是用于最大登录延迟时间
		refreshKey := fmt.Sprintf(login_key.AsmRefresh, token)
		refreshUser := &vo.RefreshUserVo{
			Uid:   userEntity.ID,
			Uname: userEntity.Name,
		}
		jsonStr, code := utils_tool.Struct2Json(refreshUser)
		if code == constapicode.Success {
			cache.Set(refreshKey, jsonStr, time.Duration(login_key.RefreshHour)*time.Hour)
		}

		slog.Infof("[login][request][/login/ [post]] jwt put redis,key-> %v.", loginKey)
	}
	// 事务提交
	dao.Commit(ctx)
	return &vo.LoginVo{
		UserId: userEntity.ID,
		Name:   userEntity.Name,
		Token:  token,
	}, constapicode.Success
}

// LoginOut 登出
func (loginApp *LoginApp) LoginOut(c *gin.Context) constapicode.SocError {
	// 获取token信息
	var token = c.Request.Header.Get(login_key.AsmToken)
	claims, code := getUserIdByJwtToken(token)
	if code != constapicode.Success {
		return code
	}

	// login key
	loginKey := fmt.Sprintf(login_key.AsmLogin, claims.UserId)
	isExist := cache.Exists(loginKey)
	if isExist {
		cache.Delete(loginKey)
	}
	//refresh key
	refreshKey := fmt.Sprintf(login_key.AsmRefresh, token)
	isExistRefreshKey := cache.Exists(refreshKey)
	if isExistRefreshKey {
		cache.Delete(refreshKey)
	}
	return constapicode.Success
}

// Refresh 用户token续期
func (loginApp *LoginApp) Refresh(c *gin.Context) (*vo.LoginVo, constapicode.SocError) {
	// 获取token信息
	var token = c.Request.Header.Get(login_key.AsmToken)
	claims, code := getUserIdByJwtToken(token)
	if code != constapicode.Success {
		slog.Errorf("refresh token[%v] is null", token)
	}
	// 判断token是否存在用户信息，如果存在说明未过期，直接返回；
	if claims != nil && claims.UserId != 0 {
		loginKey := fmt.Sprintf(login_key.AsmLogin, claims.UserId)
		isExists := cache.Exists(loginKey)
		if isExists {
			return &vo.LoginVo{
				UserId: claims.UserId,
				Name:   claims.Username,
				Token:  token,
			}, constapicode.Success
		}
	}
	// 如果token过期；则判断refresh key是否存在
	var refreshUser *vo.RefreshUserVo
	refreshKey := fmt.Sprintf(login_key.AsmRefresh, token)
	isExistRefreshKey := cache.Exists(refreshKey)
	if isExistRefreshKey {
		user, err := cache.Get(refreshKey)
		if err != nil && user != "" {
			refreshUser = getUserByRefreshToken(user)
			// 生成jwt token
			token, code = jwt_tool.GeneratorToken(refreshUser.Uid, refreshUser.Uname)
			if code != constapicode.Success {
				return nil, code
			} else {
				loginKey := fmt.Sprintf(login_key.AsmLogin, refreshUser.Uid)
				cache.Set(loginKey, token, time.Duration(login_key.LoginHour)*time.Hour)
			}
		}
	} else {
		// 最大有效期过期，则重新登录；
		return nil, constapicode.PleaseRepeatLogin
	}

	return &vo.LoginVo{
		UserId: refreshUser.Uid,
		Name:   refreshUser.Uname,
		Token:  token,
	}, constapicode.Success
}
