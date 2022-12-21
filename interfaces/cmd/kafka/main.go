package main

import (
	kfk "asm_platform/infrastructure/pkg/database/kafka"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"fmt"
)

func main() {
	consumer := kfk.KfkConsumer{
		Topic:           "asm_kafka_test1",
		EnablePartition: false,
		Partition:       0,
		GroupID:         "asm_group2",
	}
	r := consumer.NewKfkConsumer()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			slog.Errorf("read kafka message error %v", err.Error())
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		//if err := r.Close(); err != nil {
		//	slog.Errorf("failed to close reader: %v", err)
		//}
	}
}
