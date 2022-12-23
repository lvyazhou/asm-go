package contansassetcode

// AssetLevel 资产级别常量
type AssetLevel int8

const (
	// AssetLevelH 高
	AssetLevelH AssetLevel = 1

	// AssetLevelM 中
	AssetLevelM AssetLevel = 2

	// AssetLevelL 低
	AssetLevelL AssetLevel = 3
)

func (e AssetLevel) String() string {
	switch e {
	case AssetLevelH:
		return "高"
	case AssetLevelM:
		return "中"
	case AssetLevelL:
		return "低"
	default:
		return "未知"
	}
}

func (e AssetLevel) Value() int8 {
	switch e {
	case AssetLevelH:
		return 1
	case AssetLevelM:
		return 2
	case AssetLevelL:
		return 3
	default:
		return -1
	}
}
