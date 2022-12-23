package repository

import (
	"asm_platform/domain/entity/user"
)

type UserRepository interface {

	// SaveUser 用户保存
	SaveUser(user *user_entity.User) (*user_entity.User, error)

	// GetUser 获取用户
	GetUser(int64) (*user_entity.User, error)

	// FindUserList 查询用户分页
	FindUserList(userQuery *user_entity.UserQuery) ([]*user_entity.User, int64, error)

	// GetUserByAccount 通过用户邮箱查询用户
	GetUserByAccount(user *user_entity.User) (*user_entity.User, error)
}
