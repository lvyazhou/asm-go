package attack_app

import "asm_platform/application/dto"

type VncConn struct {
}

func (v *VncConn) Attack(info *dto.HostInfo, user string, pass string, timeout int64) (flag bool, err error) {
	flag = false

	return flag, nil
}
