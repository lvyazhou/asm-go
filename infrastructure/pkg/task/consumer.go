package task

import "asm_platform/infrastructure/repo"

// 读取kafka
func readKafkaMessage() {
	r := repo.NewKafkaRepo("lyz", "lyz_group_666")
	r.ReadKafkaMessageByGroupId()
}
