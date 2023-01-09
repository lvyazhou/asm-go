package attack_app

import (
	"asm_platform/application/dto"
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

type FtpConn struct {
}

func (f *FtpConn) Attack(info *dto.HostInfo, user string, pass string, timeout int64) (flag bool, err error) {
	flag = false
	Host, Port, Username, Password := info.Host, info.Port, user, pass
	conn, err := ftp.Dial(fmt.Sprintf("%v:%v", Host, Port), ftp.DialWithTimeout(time.Duration(timeout)*time.Second))
	if err == nil {
		err = conn.Login(Username, Password)
		if err == nil {
			flag = true
		} else {
			return flag, err
		}
	}
	return flag, nil
}
