package domain_entity

import (
	utils_tool "asm_platform/infrastructure/pkg/tool/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DomainQuery 域名query
type DomainQuery struct {
	Condition struct {
		// 资产ID
		AssetId string `json:"assetId"`
		// 域名
		Domain string `json:"assetName"`
	} `json:"condition"`
	// 当前页码
	PageNo int64 `json:"pageNo"`
	// 每页多少条
	PageSize int64 `json:"pageSize"`
}

func (query *DomainQuery) BuildQueryFilter() bson.M {
	filterM := bson.M{}
	filterData := make(map[string]interface{})

	// ID编号 - 完全匹配
	if query.Condition.AssetId != "" {
		id, _ := primitive.ObjectIDFromHex(query.Condition.AssetId)
		filterData["asset"] = id
	}

	// 判断域名 - 模糊匹配
	if query.Condition.Domain != "" {
		filterData["domain"] = bson.D{{"$regex", query.Condition.Domain}}
	}
	filterM = utils_tool.ConvertFilter(filterData).(bson.M)

	fmt.Println(filterM)

	return filterM
}
