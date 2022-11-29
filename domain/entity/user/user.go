package user_entity

import (
	user_vo "asm_platform/application/vo"
	"asm_platform/domain/entity"
	constuserstatus "asm_platform/infrastructure/pkg/constants/user_status"
	utils_tool "asm_platform/infrastructure/pkg/tool/utils"
	"reflect"
)

// User 用户实体
type User struct {
	entity.BasicEntity

	// Account 账户名称
	Account string `gorm:"column:u_account"`

	// Email 电子邮件
	Email string `gorm:"column:email"`

	// Name 姓名
	Name string `gorm:"column:u_name"`

	// Mobile 手机号码
	Mobile string `gorm:"column:u_mobile"`

	// Password 密码
	Password string `gorm:"column:u_pwd"`

	// Status 用户状态 @see constants.user_status
	Status int8 `gorm:"column:u_status"`
}

func (User) TableName() string {
	return "u_user"
}

// IsEmpty 判断是否为空
func (user *User) IsEmpty() bool {
	if user == nil {
		return false
	}
	return reflect.DeepEqual(user, &User{})
}

// UserEntityToVo 转化实体 entity - vo
func (user User) UserEntityToVo() *user_vo.UserVO {
	return &user_vo.UserVO{
		BasicVo: user_vo.BasicVo{
			ID:         user.ID,
			CreateTime: utils_tool.FormatTimeToString(user.CreateTime),
			UpdateTime: utils_tool.FormatTimeToString(user.UpdateTime),
		},
		Account:  user.Account,
		Name:     user.Name,
		Email:    user.Email,
		Mobile:   user.Mobile,
		Status:   constuserstatus.UserStatus(user.Status).String(),
		RoleList: nil,
	}
}
