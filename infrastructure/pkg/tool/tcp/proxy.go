// @Title  tcp链接工具包
// @Description:
// @Author: lvyazhou
// @Date: 2022/7/4 16:12

package tcp_tools

import (
	"net"
	"time"
)

func WrapperTcpWithTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := &net.Dialer{Timeout: timeout}
	return WrapperTCP(network, address, d)
}

func WrapperTCP(network, address string, forward *net.Dialer) (net.Conn, error) {
	//get conn
	var conn net.Conn
	var err error
	conn, err = forward.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return conn, nil

}
