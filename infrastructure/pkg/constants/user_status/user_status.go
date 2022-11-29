package constuserstatus

// UserStatus 用户状态常量
type UserStatus int8

const (
	// UserStatusNormal 正常/启用
	UserStatusNormal UserStatus = 1

	// UserStatusDisable 禁用
	UserStatusDisable UserStatus = 2

	// UserStatusDeleted 删除
	UserStatusDeleted UserStatus = 3

	// UserStatusUnknown 未知
	UserStatusUnknown UserStatus = -1
)

func (e UserStatus) String() string {
	switch e {
	case UserStatusNormal:
		return "正常"
	case UserStatusDisable:
		return "已禁用"
	case UserStatusDeleted:
		return "已删除"
	default:
		return "未知"
	}
}

func (e UserStatus) Value() int8 {
	switch e {
	case UserStatusNormal:
		return 1
	case UserStatusDisable:
		return 2
	case UserStatusDeleted:
		return 3
	default:
		return -1
	}
}
