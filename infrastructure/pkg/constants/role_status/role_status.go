package constrolestatus

// RoleStatus 角色状态常量
type RoleStatus int8

const (
	// RoleStatusNormal 正常/启用
	RoleStatusNormal RoleStatus = 1

	// RoleStatusDisable 禁用
	RoleStatusDisable RoleStatus = 2

	// RoleStatusDeleted 删除
	RoleStatusDeleted RoleStatus = 3
)

func (e RoleStatus) String() string {
	switch e {
	case RoleStatusNormal:
		return "正常"
	case RoleStatusDisable:
		return "已禁用"
	case RoleStatusDeleted:
		return "已删除"
	default:
		return "未知"
	}
}

func (e RoleStatus) Val() (val int8) {
	switch e {
	case RoleStatusNormal:
		val = 1
	case RoleStatusDisable:
		val = 2
	case RoleStatusDeleted:
		val = 3
	default:

	}

	return
}
