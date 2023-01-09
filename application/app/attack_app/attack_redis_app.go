package attack_app

import (
	"asm_platform/application/dto"
	tcp_tools "asm_platform/infrastructure/pkg/tool/tcp"
	"fmt"
	"net"
	"strings"
	"time"
)

type RedisConn struct {
}

func (m *RedisConn) Attack(info *dto.HostInfo, user string, pass string, timeout int64) (flag bool, err error) {
	flag, err = redisConnAuth(info, pass, timeout)
	return flag, err
}

// RedisConnAuth 用户密码redis认证
func redisConnAuth(info *dto.HostInfo, pass string, timeout int64) (flag bool, err error) {
	flag = false
	realHost := fmt.Sprintf("%s:%v", info.Host, info.Port)
	conn, err := tcp_tools.WrapperTcpWithTimeout("tcp", realHost, time.Duration(timeout)*time.Second)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err != nil {
		return flag, err
	}
	err = conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	if err != nil {
		return flag, err
	}
	_, err = conn.Write([]byte(fmt.Sprintf("auth %s\r\n", pass)))
	if err != nil {
		return flag, err
	}
	reply, err := readReply(conn)
	if err != nil {
		return flag, err
	}
	if strings.Contains(reply, "+OK") {
		flag = true
	}
	return flag, err
}

func readReply(conn net.Conn) (result string, err error) {
	size := 5 * 1024
	buf := make([]byte, size)
	for {
		count, err := conn.Read(buf)
		if err != nil {
			break
		}
		result += string(buf[0:count])
		if count < size {
			break
		}
	}
	return result, err
}
