package handler

import (
	"asm_platform/infrastructure/pkg/constants/api_code"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response api响应公共字段
type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponsePage 分页api响应公共字段
type ResponsePage struct {
	Code    int32     `json:"code"`
	Message string    `json:"message"`
	Data    *PageData `json:"data"`
}

type PageData struct {
	TotalCount int64       `json:"total"`
	List       interface{} `json:"items"`
}

// DicKVData 字典下拉响应体
type DicKVData struct {
	// id
	ID int64 `json:"id,string"`

	// 字典value
	DictionaryValue string `json:"dictionary_value"`
}

// ReturnFormat 简化返回格式，减少重复代码
func ReturnFormat(c *gin.Context, code constapicode.SocError, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code.Val(),
		Message: code.String(),
		Data:    data,
	})
}

// ReturnPageFormat 简化返回格式，减少重复代码
func ReturnPageFormat(c *gin.Context, code constapicode.SocError, data interface{}, totalCount int64) {
	c.JSON(http.StatusOK, ResponsePage{
		Code:    code.Val(),
		Message: code.String(),
		Data: &PageData{
			TotalCount: totalCount,
			List:       data,
		},
	})
}
