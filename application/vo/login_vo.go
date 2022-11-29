package vo

import "reflect"

// LoginVo 登录返回VO
type LoginVo struct {
	// 用户ID
	UserId int64 `json:"user_id"`
	// 用户名称
	Name string `json:"name"`
	// 用户token
	Token string `json:"token"`
}

// RefreshUserVo 用户续期token
type RefreshUserVo struct {
	// 用户ID
	Uid int64 `json:"uid"`
	// 用户名称
	Uname string `json:"uname"`
}

func (user RefreshUserVo) IsEmpty() bool {
	return reflect.DeepEqual(user, RefreshUserVo{})
}

func (user LoginVo) IsEmpty() bool {
	return reflect.DeepEqual(user, LoginVo{})
}
