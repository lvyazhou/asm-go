package oplog_app

import (
	"asm_platform/application/dto"
	"asm_platform/application/vo"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
)

type OpLogAppInterface interface {
	// SaveOpLog 保存操作日志
	SaveOpLog(oplog *dto.OpLogDTO) constapicode.SocError

	// EditOpLog 更新操作日志
	EditOpLog(oplog *dto.OpLogDTO) constapicode.SocError

	// SearchOpLogList 查询日志分页
	SearchOpLogList(opLogQuery *dto.OpLogQueryDTO) ([]vo.OpLogVo, int64, constapicode.SocError)

	// DeleteOpLogById 删除操作日志 根据ID
	DeleteOpLogById(id string) constapicode.SocError

	// DeleteOpLogByQuery 删除操作日志 根据查询条件
	DeleteOpLogByQuery(opLogQuery *dto.OpLogQueryDTO) constapicode.SocError
}
