package user_entity

import (
	"asm_platform/domain/entity"
	"reflect"
)

// UserQuery 用户查询实体
type UserQuery struct {
	// 公共实体
	entity.BasicQueryEntity
	// 用户名
	Uname string `json:"uname"`
	// 邮箱
	Email string `json:"email"`
}

// IsEmpty 判断是否为空
func (userQuery UserQuery) IsEmpty() bool {
	return reflect.DeepEqual(userQuery, User{})
}
