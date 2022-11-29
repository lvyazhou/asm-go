package handler

// 分页请求公共字段
type RequestPage struct {

	// 分页页面
	Page int `json:"page" binding:"required,min=0"`

	// 分页大小 默认10
	Size int `json:"pageSize" binding:"required,min=0"`
}
