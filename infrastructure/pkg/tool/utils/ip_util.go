// @Title  IP工具类
// @Description:
// @Author: lvyazhou
// @Date: 2022/7/25 14:33

package utils_tool

import "net"

//
//  判断是否是公网ip
//  @Description:
//  @params IP   ip
//  @return bool true:公网 false:内网
//
func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}
