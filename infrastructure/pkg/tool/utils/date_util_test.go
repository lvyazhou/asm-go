// @Title  日期工具测试
// @Description:
// @Author: lvyazhou
// @Date: 2022/6/17 16:11

package utils_tool

import (
	"fmt"
	"testing"
)

// 测试 日期间隔内的所有日期
func TestGetBetweenDates(t *testing.T) {
	dates := GetBetweenDates("2022-06-29", "2022-07-03")
	fmt.Print(dates)
}
