// @Title  业务工具
// @Description:
// @Author: lvyazhou
// @Date: 2022/7/4 16:12

package utils_tool

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"
)

// GenerateBllSn
// 生成编号（sn）
// @Description: 前缀 + 时间 + 5位随机数
// @params prefix
//
func GenerateBllSn(prefix string, num int) string {
	// 时间处理
	t1 := strings.ReplaceAll(Timestamp13LongDateStr(time.Now().UnixMilli()), "-", "")
	t2 := strings.ReplaceAll(t1, " ", "")
	t3 := strings.ReplaceAll(t2, ":", "")

	var randomStr = ""
	// 生成5位随机数
	for i := 0; i < num; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(9))
		randomStr = randomStr + fmt.Sprint(r)
	}

	return fmt.Sprint(prefix, t3, randomStr)
}
