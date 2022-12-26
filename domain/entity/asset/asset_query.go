package asset_entity

import utils_tool "asm_platform/infrastructure/pkg/tool/utils"

// AssetQuery 资产查询
type AssetQuery struct {
	Condition struct {
		// 资产名称
		AssetName string `json:"assetName"`

		// 资产类型
		AssetType string `json:"assetType"`

		// 资产发现时间
		StartTime string `json:"startTime"`
		EndTime   string `json:"endTime"`
	} `json:"condition"`
	// 当前页码
	PageNo int64 `json:"pageNo"`
	// 每页多少条
	PageSize int64 `json:"pageSize"`
}

// query to filter

func (query *AssetQuery) BuildAssetQueryFilter() interface{} {
	filterData := make(map[string]interface{})
	// 判断资产名称 - 模糊匹配
	if query.Condition.AssetName != "" {
		filterName := make(map[string]string)
		filterName["$regex"] = query.Condition.AssetName
		filterData["asset_name"] = filterName
	}
	// 资产类型 - 完全匹配
	if query.Condition.AssetType != "" {
		filterData["asset_type"] = query.Condition.AssetType
	}
	// 开始和结束时间
	if query.Condition.StartTime != "" && query.Condition.EndTime != "" {
		filterTime := make(map[string]string)
		filterTime["$gte"] = query.Condition.StartTime
		filterTime["$lte"] = query.Condition.EndTime
		filterData["create_time"] = filterTime
	}
	return utils_tool.ConvertFilter(filterData)
}
