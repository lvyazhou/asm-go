package dto

// PageCommon 分页公共实体
type PageCommon struct {
	// 分页页面 从0开始
	Page int `json:"page"`

	// 分页大小 默认10
	Size int `json:"pageSize"`
}

// KVInfo key-value
type KVInfo struct {

	// id
	Id int64 `json:"id"`

	// 名称
	Name string `json:"name"`
}

type ResponseKV struct {
	// 响应编码
	Code int `json:"code"`
	// 响应内容
	Message string `json:"message"`
}
