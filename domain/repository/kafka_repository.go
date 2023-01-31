package repository

type KafkaRepository interface {
	// WriteKafkaMessage 写入kafka message
	WriteKafkaMessage([]map[string]string) error
	// ReadKafkaMessage 读取kafka message
	ReadKafkaMessage()

	ReadKafkaMessageByGroupId()
}
