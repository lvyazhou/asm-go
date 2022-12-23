package asset_entity

// AssetQuery 资产查询
type AssetQuery struct {

	// 资产名称
	AssetName string `json:"assetName"`

	// 当前页码
	PageNo int64 `json:"pageNo"`
	// 每页多少条
	PageSize int64 `json:"pageSize"`
}
