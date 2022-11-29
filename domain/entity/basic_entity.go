package entity

import "time"

// BasicEntity 数据基本字段
type BasicEntity struct {

	// ID 雪花算法生成
	ID int64 `gorm:"primaryKey"`

	// CreateUser 创建人ID
	CreateUser int64 `gorm:"column:create_user_id"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"column:create_time"`

	// UpdateUser 更新人ID
	UpdateUser int64 `gorm:"column:update_user_id"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"column:update_time"`
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
