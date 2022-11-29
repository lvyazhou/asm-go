package mds

import (
	"asm_platform/application/dto"
	"asm_platform/infrastructure/pkg/constants"
	"asm_platform/infrastructure/pkg/slog"
	"asm_platform/infrastructure/pkg/task"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func GinLogger(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求参数
		args := GetRequestArgs(c)
		// 请求url
		path := c.Request.URL.Path
		// 客户端IP
		clientIp := c.ClientIP()
		// 响应体
		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		//logger.Debugf("url=%s, request=%s,status=%d, response=%s , clientIP=%s", path, args, c.Writer.Status(), blw.body.String(), clientIp)

		// 放入log 任务队列中
		var userId int64
		var userName string

		// 过滤验证码和登录api接口
		if !constants.FilterUserApiPath.Contains(path) {
			if _, exists := c.Get(constants.UserId); exists {
				userId = c.MustGet(constants.UserId).(int64)
				userName = c.MustGet(constants.UserName).(string)
			}
		}
		// 过滤验证码和审计日志不产生日志
		if !constants.FilterLogApiPath.Contains(path) {
			result := dto.OpLogDTO{
				RequestArgs:    args,
				RequestUri:     path,
				ResponseResult: blw.body.String(),
				ClientIp:       clientIp,
				TimeStamp:      time.Now().UnixMilli(),
				UserId:         userId,
				UserName:       userName,
			}
			task.InsertionOpLogQueue <- result
		}
	}
}
func GinRecovery(logger *zap.SugaredLogger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// GetRequestArgs 获取请求参数，格式化为字符串
func GetRequestArgs(c *gin.Context) (args string) {
	var contentType = c.ContentType()
	//fmt.Println("request content type: ", contentType)
	if contentType == "application/json" {
		args = bodyToJson(c)
	} else if contentType == "multipart/form-data" {
		args = paramToJson(c)
	} else {
		args = c.Request.URL.RawQuery
	}
	return args
}

// form param 格式化json
func paramToJson(c *gin.Context) string {
	c.Request.ParseMultipartForm(128) //保存表单缓存的内存大小128M
	data := c.Request.Form
	if len(data) > 0 {
		jsonStr, err := json.Marshal(data)
		if err == nil {
			return string(jsonStr)
		}
	}
	return ""
}

// body to json
func bodyToJson(c *gin.Context) string {
	// 多次绑定requestbody问题
	data, err := c.GetRawData()
	if err != nil {
		slog.Errorf("body to json ge raw data error %v", err.Error())
		return ""
	}
	//fmt.Printf("data: %v\n", string(data))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 关键点
	return string(data)
}
