// @Title  Title
// @Description:
// @Author: lvyazhou
// @Date: 2022/7/25 14:38

package utils_tool

import (
	"net"
	"testing"
)

func TestIsPublicIP(t *testing.T) {
	ip := net.ParseIP("192.168.1.10")
	isPublicIP := IsPublicIP(ip)
	t.Logf("%v", isPublicIP)
}
