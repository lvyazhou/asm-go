package dto

import (
	"asm_platform/domain/entity"
	oplog_entity "asm_platform/domain/entity/oplog"
)

type OpLogDTO struct {
	// ID
	ID int64 `json:"id"`

	// 请求参数
	RequestArgs string `json:"request_args"`

	// 请求url
	RequestUri string `json:"request_uri"`

	// 响应结果
	ResponseResult string `json:"response_result"`
	ResponseKV

	// 客户端IP
	ClientIp string `json:"client_ip"`

	// 当前时间
	TimeStamp int64 `json:"timestamp"`

	//操作用户
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

// OpLogDtoTOEntity dto to entity
func (dto OpLogDTO) OpLogDtoTOEntity() *oplog_entity.OpLog {
	opLogInfo := &oplog_entity.OpLog{
		ID:          dto.ID,
		RequestUri:  dto.RequestUri,
		RequestArgs: dto.RequestArgs,
		ResponseResult: entity.ResponseKV{
			Code:    dto.Code,
			Message: dto.Message,
		},
		Timestamp: dto.TimeStamp,
		User: entity.KVInfo{
			Id:   dto.UserId,
			Name: dto.UserName,
		},
		ClientIp: dto.ClientIp,
	}
	return opLogInfo
}

type OpLogQueryDTO struct {
	PageCommon
	// 请求url
	RequestUri string `json:"request_uri"`
}

// OpLogQueryDtoTOEntity query dto to entity
func (opLogQuery OpLogQueryDTO) OpLogQueryDtoTOEntity() *oplog_entity.OpLogQuery {
	return &oplog_entity.OpLogQuery{
		BasicQueryEntity: entity.BasicQueryEntity{
			Page: opLogQuery.Page,
			Size: opLogQuery.Size,
		},
		RequestUri: opLogQuery.RequestUri,
	}
}
