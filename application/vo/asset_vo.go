package vo

import "reflect"

type AssetVo struct {
	// 资产编号
	ID string `json:"string"`
	//资产标识
	AssetNo string `json:"asset_no"`
	//资产名称
	AssetName string `json:"asset_name"`
	// 资产类型
	AssetType string `json:"asset_type"`
	// 资产级别
	AssetLevel string `json:"asset_level"`
	// 资产组
	AssetGroup string `json:"asset_group"`
	// 资产负责人
	Mgr string `json:"mgr"`
	// 资产部门
	Dept string `json:"dept"`
	// 创建人员
	CreateUser string `json:"create_user"`
	// 创建时间
	CreateTime string `json:"create_time"`
}

func (vo AssetVo) IsEmpty() bool {
	return reflect.DeepEqual(vo, AssetVo{})
}
