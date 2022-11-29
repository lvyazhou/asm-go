package vo

import (
	"reflect"
)

type OpLogVo struct {
	// id
	ID int64 `json:"id"`

	// 请求url
	RequestUri string `json:"request_uri"`

	// 请求参数
	RequestArgs string `json:"request_args"`

	// 响应结果
	ResponseResult ResponseKV `json:"response"`

	// 操作时间
	TimestampStr string `json:"timestamp_str"`

	// 操作人员
	User string `json:"user"`

	// 客户端IP
	ClientIp string `json:"client_ip"`
}

func (vo OpLogVo) IsEmpty() bool {
	return reflect.DeepEqual(vo, OpLogVo{})
}
