package oplog_handle

import (
	"asm_platform/application/app/oplog_app"
	oplog_dto "asm_platform/application/dto"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	"asm_platform/infrastructure/pkg/slog"
	"asm_platform/interfaces/handler"
	"github.com/gin-gonic/gin"
)

type OpLogs struct {
	oplog oplog_app.OpLogAppInterface
}

func NewOpLogs(oplog oplog_app.OpLogAppInterface) *OpLogs {
	return &OpLogs{oplog: oplog}
}

// ListOpLogs    godoc
// @Summary      查询审计日志
// @Description  查询审计日志
// @Tags         日志管理
// @Accept       json
// @Param        requestUri  path      string  true  "请求uri"
// @Success      200  {object}  handler.Response
// @Router       /oplog/list/ [post]
// @Security     ApiKeyAuth
func (o *OpLogs) ListOpLogs(c *gin.Context) {
	var param = &oplog_dto.OpLogQueryDTO{}
	if err := c.ShouldBindJSON(param); err != nil {
		slog.Errorf("[oplog][request][/oplog/list/ [post]] listOpLogs valid error %v.", err.Error())
		handler.ReturnFormat(c, constapicode.ErrorReq, nil)
		return
	}
	oplogList, totalCount, code := o.oplog.SearchOpLogList(param)
	handler.ReturnPageFormat(c, code, oplogList, totalCount)
	return
}
