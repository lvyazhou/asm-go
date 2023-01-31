package attack_app

import "asm_platform/application/dto"

type EsConn struct {
}

func (e *EsConn) Attack(info *dto.HostInfo, user string, pass string, timeout int64) (flag bool, err error) {
	flag = false

	return flag, nil
}
