package task

import (
	"asm_platform/application/dto"
)

var (
	InsertionOpLogQueue chan dto.OpLogDTO
)

// Init 任务队列初始化
func Init() {
	// request op log task queue
	InsertionOpLogQueue = make(chan dto.OpLogDTO, 10000)
	go InsertTaskQueueResult()

	// .........other

	go readKafkaMessage()
}
