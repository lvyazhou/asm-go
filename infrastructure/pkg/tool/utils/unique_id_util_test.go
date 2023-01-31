package utils_tool

import "testing"

func TestGenerateUid(t *testing.T) {
	t.Run("生成雪花算法ID", func(t *testing.T) {
		sn := GenerateUniqueUint64Id()
		t.Log(sn)
	})
}
