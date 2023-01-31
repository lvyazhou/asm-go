package attack_app

import (
	"asm_platform/application/dto"
	constservicetype "asm_platform/infrastructure/pkg/constants/service_type"
	"testing"
)

func TestGetBruteType(t *testing.T) {
	// test mysql 4
	st := constservicetype.ServiceType(1)

	info := &dto.HostInfo{
		Host: "127.0.0.1",
		Port: "3308",
	}

	var Passwords = []string{
		"123456", "admin", "admin123", "root", "", "pass123", "pass@123", "password", "123123", "654321", "111111", "123", "1", "admin@123", "Admin@123", "admin123!@#", "{user}", "{user}1", "{user}111", "{user}123", "{user}@123", "{user}_123", "{user}#123", "{user}@111", "{user}@2019", "{user}@123#4", "P@ssw0rd!", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe", "123qwe!@#", "123456789", "123321", "666666",
		"a123456.", "123456~a", "123456!a", "000000", "1234567890", "8888888", "!QAZ2wsx",
		"1qaz2wsx", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345",
		"a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.", "sysadmin", "system", "1qaz!QAZ", "2wsx@WSX",
		"qwe123!@#", "Aa123456!", "A123456s!", "sa123456", "1q2w3e",
		"xmap",
		"Charge123",
		"Aa123456789",
	}
	var Users = []string{"root", "xmap"}

	cf := GetBruteConfig(st, info, 10)
	infos := cf.GenerateData(Users, Passwords)
	ret := cf.Run(100, infos)
	t.Log(ret)
	t.Log("爆破次数", cf.GetNum())
	if ret {

		t.Log(cf.GetInfo())
	}
}
