package attack_app

import (
	"asm_platform/application/dto"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

type SshConn struct {
}

func (s *SshConn) Attack(info *dto.HostInfo, user string, pass string, timeout int64) (flag bool, err error) {
	flag = false
	Host, Port, Username, Password := info.Host, info.Port, user, pass
	Auth := []ssh.AuthMethod{ssh.Password(Password)}

	config := &ssh.ClientConfig{
		User:    Username,
		Auth:    Auth,
		Timeout: time.Duration(timeout) * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", Host, Port), config)
	if err == nil {
		defer client.Close()
		session, err := client.NewSession()
		if err == nil {
			defer session.Close()
			flag = true
		}
	}
	return flag, nil
}
