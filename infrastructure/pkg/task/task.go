package task

import (
	"asm_platform/application/app/oplog_app"
)

// InsertTaskQueueResult 任务队列处理
func InsertTaskQueueResult() {
	for {
		select {
		// 如果是请求日志
		case v := <-InsertionOpLogQueue:
			oplogApp := oplog_app.NewOpLogApp()
			go oplogApp.SaveOpLog(&v)
			//fmt.Println("soc op log task queue ", v)
		}
		// 如果是其他......
	}
}
