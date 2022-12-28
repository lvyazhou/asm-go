package domain_handle

import (
	"asm_platform/application/app/domain_app"
	"asm_platform/infrastructure/pkg/constants/api_code"
	"asm_platform/interfaces/handler"
	"github.com/gin-gonic/gin"
)

type DomainHandle struct {
	dh domain_app.DomainAppInterface
}

func NewDomainHandle() *DomainHandle {
	return &DomainHandle{dh: domain_app.NewDomainApp()}
}

func (dh *DomainHandle) SaveDomain(c *gin.Context) {
	dh.dh.SaveDomain()
	handler.ReturnFormat(c, constapicode.Success, nil)
	return
}

func (dh *DomainHandle) FindDomainList(c *gin.Context) {
	list, code := dh.dh.FindDomainList()
	handler.ReturnFormat(c, code, list)
	return
}
