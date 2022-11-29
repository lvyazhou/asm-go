package dto

import (
	"asm_platform/domain/entity"
	user_entity "asm_platform/domain/entity/user"
	slat_tool "asm_platform/infrastructure/pkg/tool/slat"
	utils_tool "asm_platform/infrastructure/pkg/tool/utils"
	"time"
)

// UserDTO 用户请求DTO
type UserDTO struct {
	// Account 账户名称
	Account string `json:"account" binding:"required,min=4,max=32"`

	// 密码
	PassWord string `json:"pass_word"`

	// Email 电子邮件
	Email string `json:"email" binding:"required,min=1,max=64"`

	// Name 姓名
	Name string `json:"name" binding:"required,min=1,max=32"`

	// Mobile 手机号码
	Mobile string `json:"mobile" binding:"required,len=11"`

	// Status 用户状态 @see constants.user_status 1:正常；2：禁用；3：已删除'
	Status int8 `json:"status" binding:"min=0,max=3"`

	// 角色列表
	RoleList []string `json:"role_list"`
}

// UserQueryDto 用户查询DTO
type UserQueryDto struct {
	PageCommon
	// 用户名
	Uname string `json:"uname"`
	// 邮箱
	Email string `json:"email"`
}

// UserStatusDto 用户状态DTO
type UserStatusDto struct {
	// 用户ID
	ID string `json:"id"`
	// Status 用户状态 @see constants.user_status 1:正常；2：禁用；3：已删除'
	Status int8 `json:"status" binding:"min=0,max=3"`
}

// UserPwdDto 用户密码DTO
type UserPwdDto struct {
	// 用户ID
	ID string `json:"id"`

	// OldPassword 旧密码
	OldPassword string `json:"old_password" binding:"required,min=8,max=32"`

	// NewPassword 新密码
	NewPassword string `json:"new_password" binding:"required,min=8,max=32"`
}

// UserDtoToEntity 转化实体 dto - entity
func (dto UserDTO) UserDtoToEntity(id, uid int64) *user_entity.User {
	return &user_entity.User{
		BasicEntity: entity.BasicEntity{
			ID:         id,
			CreateUser: uid,
			CreateTime: time.Now(),
		},
		Account:  dto.Account,
		Email:    dto.Email,
		Name:     dto.Name,
		Mobile:   dto.Mobile,
		Password: utils_tool.Md5Crypto(slat_tool.SlatUserPwd + dto.PassWord),
		Status:   dto.Status,
	}
}

// UserQueryDtoToEntity 转化实体 dto - entity
func (dto UserQueryDto) UserQueryDtoToEntity() *user_entity.UserQuery {
	return &user_entity.UserQuery{
		BasicQueryEntity: entity.BasicQueryEntity{
			Page: dto.Page,
			Size: dto.Size,
		},
		Uname: dto.Uname,
		Email: dto.Email,
	}
}
