package attack_app

import "asm_platform/application/dto"

type BruteFunc interface {
	Attack(info *dto.HostInfo, user string, pass string, timeout int64) (flag bool, err error)
}
