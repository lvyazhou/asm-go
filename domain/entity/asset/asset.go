package asset_entity

import "time"

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
