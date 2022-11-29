package oplog_entity

import (
	"asm_platform/domain/entity"
	"reflect"
)

// OpLogQuery 日志查询实体
type OpLogQuery struct {
	// 公共实体
	entity.BasicQueryEntity

	// 请求URL
	RequestUri string `json:"request_uri"`

	// ID
	ID int64 `json:"id"`
}

// IsEmpty 判断是否为空
func (query OpLogQuery) IsEmpty() bool {
	return reflect.DeepEqual(query, OpLogQuery{})
}
