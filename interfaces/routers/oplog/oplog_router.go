package oplog_router

import (
	"asm_platform/application/app/oplog_app"
	oplog_handle "asm_platform/interfaces/handler/oplog"
	"github.com/gin-gonic/gin"
)

func SetupOpLogRouter(group *gin.RouterGroup) {
	oplogGroup := group.Group("/oplog")
	{
		// 实例化应用层接口
		oplogs := oplog_handle.NewOpLogs(oplog_app.NewOpLogApp())

		// 审计日志查询分页
		oplogGroup.POST("/list/", oplogs.ListOpLogs)

	}
}
