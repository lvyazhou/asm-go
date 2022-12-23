package repository

import (
	"asm_platform/domain/entity/oplog"
)

type OpLogRepository interface {
	// SaveOpLog 保存操作日志
	SaveOpLog(opLog *oplog_entity.OpLog) error

	//BatchSaveOpLog 批量保存操作日志
	BatchSaveOpLog(opLogMap map[string]string) error

	// FindOpLogList 日志查询分页
	FindOpLogList(query *oplog_entity.OpLogQuery) ([]oplog_entity.OpLog, int64, error)

	// DeleteOpLogById 根据ID删除日志
	DeleteOpLogById(id string) error

	// DeleteOpLogByQuery 根据query删除日志
	DeleteOpLogByQuery(query *oplog_entity.OpLogQuery) (int64, error)

	// DeleteOpLogByIds 根据IDs批量删除日志
	DeleteOpLogByIds(ids []string) (int, error)
}
