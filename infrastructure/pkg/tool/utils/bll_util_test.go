// @Title  Title
// @Description:
// @Author: lvyazhou
// @Date: 2022/7/4 16:18

package utils_tool

import "testing"

func TestGenerateBllSn(t *testing.T) {
	t.Run("生成sn", func(t *testing.T) {
		sn := GenerateBllSn("FW", 5)
		t.Log("sn：" + sn)
	})
}
