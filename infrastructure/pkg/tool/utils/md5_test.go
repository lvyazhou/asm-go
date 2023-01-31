package utils_tool

import (
	slat2 "asm_platform/infrastructure/pkg/tool/slat"
	"fmt"
	"testing"
)

func TestMd5Crypto(t *testing.T) {
	str := "123456789011"
	crypto := Md5Crypto(slat2.SlatUserPwd + str)
	fmt.Println(crypto)
}
