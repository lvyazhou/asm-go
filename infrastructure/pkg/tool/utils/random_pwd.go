package utils_tool

import (
	"crypto/rand"
	"fmt"
)

//
// GetRandomPWD
// @Description: 生成随机密码
// @params n 长度
// @return string
//
func GetRandomPWD(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
