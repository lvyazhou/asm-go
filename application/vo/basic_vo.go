package vo

// vo 公共
type BasicVo struct {
	// ID 雪花算法生成
	ID int64 `json:"id"`

	// CreateUser 创建人ID
	CreateUser string `json:"create_user"`

	// CreateTime 创建时间
	CreateTime string `json:"create_time"`

	// UpdateUser 更新人ID
	UpdateUser string `json:"update_user"`

	// UpdateTime 更新时间
	UpdateTime string `json:"update_time"`
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
