package repository

type KafkaRepository interface {
	// WriteKafkaMessage 单条写入kafka message
	WriteKafkaMessage(map[string]string) error
	// WriteKafkaMessageList 批量写入kafka message
	WriteKafkaMessageList([]map[string]string) error
	// ReadKafkaMessage 读取kafka message
	ReadKafkaMessage()
	ReadKafkaMessageByGroupId()
}
