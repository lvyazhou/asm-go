package asset_entity

import (
	"asm_platform/application/vo"
	contansassetcode "asm_platform/infrastructure/pkg/constants/asset_code"
	utils_tool "asm_platform/infrastructure/pkg/tool/utils"
	"reflect"
	"time"
)

// Asset 网络资产属性
type Asset struct {

	// ID
	ID int64 `bson:"id"`

	// 资产标识
	AssetNo string `bson:"asset_no"`

	// 资产名称
	AssetName string `bson:"asset_name"`

	// 资产类型
	AssetType string `bson:"asset_type"`

	// 资产级别
	AssetLevel int8 `bson:"asset_level"`

	// 资产负责人
	Mgr string `bson:"mgr,omitempty"`

	// 所属部门
	Dept string `bson:"dept,omitempty"`

	// 资产组
	AssetGroup string `bson:"asset_group,omitempty"`

	// CreateUser 创建人ID
	CreateUser int64 `bson:"create_user_id,omitempty"`

	// CreateTime 创建时间
	CreateTime time.Time `bson:"create_time,omitempty"`

	// UpdateUser 更新人ID
	UpdateUser int64 `bson:"update_user_id,omitempty"`

	// UpdateTime 更新时间
	UpdateTime time.Time `bson:"update_time,omitempty"`
}

// IsEmpty 判断是否为空
func (asset *Asset) IsEmpty() bool {
	if asset == nil {
		return false
	}
	return reflect.DeepEqual(asset, Asset{})
}

// AssetToVo 实体转化为vo
func (asset *Asset) AssetToVo() *vo.AssetVo {
	return &vo.AssetVo{
		ID:         asset.ID,
		AssetNo:    asset.AssetNo,
		AssetName:  asset.AssetName,
		AssetType:  asset.AssetType,
		AssetLevel: contansassetcode.AssetLevel(asset.AssetLevel).String(),
		AssetGroup: asset.AssetGroup,
		Mgr:        asset.Mgr,
		Dept:       asset.Dept,
		CreateTime: utils_tool.FormatTimeToString(asset.CreateTime),
	}
}
