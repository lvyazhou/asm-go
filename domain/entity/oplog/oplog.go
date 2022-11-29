package oplog_entity

import (
	"asm_platform/application/vo"
	"asm_platform/domain/entity"
	utils_tool "asm_platform/infrastructure/pkg/tool/utils"
	"reflect"
)

// OpLog 操作实体类
type OpLog struct {
	// id
	ID int64 `json:"id"`

	// 请求url
	RequestUri string `json:"request_uri"`

	// 请求参数
	RequestArgs string `json:"request_args"`

	// 响应结果
	ResponseResult entity.ResponseKV `json:"response"`

	// 所属客户
	Customer entity.KVInfo `json:"customer"`

	// 操作时间
	Timestamp int64 `json:"timestamp"`

	// 操作人员
	User entity.KVInfo `json:"user"`

	// 客户端IP
	ClientIp string `json:"client_ip"`
}

// IsEmpty 判断是否为空
func (opLog OpLog) IsEmpty() bool {
	return reflect.DeepEqual(opLog, OpLog{})
}

// OpLogToVo entity to vo
func (opLog OpLog) OpLogToVo() vo.OpLogVo {
	opLogInfo := vo.OpLogVo{
		ID:          opLog.ID,
		RequestUri:  opLog.RequestUri,
		RequestArgs: opLog.RequestArgs,
		ResponseResult: vo.ResponseKV{
			Code:    opLog.ResponseResult.Code,
			Message: opLog.ResponseResult.Message,
		},
		TimestampStr: utils_tool.Timestamp13LongDateStr(opLog.Timestamp),
		User:         opLog.User.Name,
		ClientIp:     opLog.ClientIp,
	}
	return opLogInfo
}
