package entity

// BasicQueryEntity 数据查询基本字段
type BasicQueryEntity struct {
	
	// 分页页面 从0开始
	Page int `json:"page"`

	// 分页大小 默认10
	Size int `json:"pageSize"`
}
