package kfk

import (
	"asm_platform/infrastructure/config"
	"asm_platform/infrastructure/pkg/slog"
	"github.com/segmentio/kafka-go"
)

var Conn *kafka.Conn

// Init 初始化kafka链接
func Init() {
	if err := newKafkaDriver(); err != nil {
		slog.Panicf("kafka error on database initialization: %s\n", err)
		return
	}
}
func newKafkaDriver() error {
	conf := config.GetConfig()
	brokers := conf.GetString("kafka.addr")

	conn, err := kafka.Dial("tcp", brokers)
	if err != nil {
		slog.Errorf("init kafka connection error %v", err.Error())
		return err
	}
	Conn = conn
	slog.Infof("kafka Connected...")
	//defer conn.Close()
	return nil
}

// KfkProduct kafka生产者传参
type KfkProduct struct {
	Topic string `json:"topic"`
}

// NewKfkProduct kafka生产者实例生成
func (product *KfkProduct) NewKfkProduct() *kafka.Writer {
	conf := config.GetConfig()
	addr := conf.GetString("kafka.addr")
	w := &kafka.Writer{
		Addr:                   kafka.TCP(addr),
		Topic:                  product.Topic,
		Balancer:               &kafka.Hash{},
		AllowAutoTopicCreation: true,         // 主体创建，不存在时
		Compression:            kafka.Snappy, // 压缩消息体
	}
	return w
}

// KfkConsumer kafka消费者传参
type KfkConsumer struct {
	Brokers         []string `json:"brokers"`
	Topic           string   `json:"topic"`
	EnablePartition bool     `json:"enablePartition"`
	Partition       int      `json:"partition"`
	GroupID         string   `json:"groupID"`
}

// NewKfkConsumer kafka消费者实例生成 根据partition / groupId 两者只能分配一个
func (consumer *KfkConsumer) NewKfkConsumer() *kafka.Reader {
	conf := config.GetConfig()
	addr := conf.GetString("kafka.addr")
	confReader := kafka.ReaderConfig{
		Brokers:  []string{addr},
		Topic:    consumer.Topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		//CommitInterval: time.Second * 1,
	}
	// 判断是否指定partition分片
	if consumer.EnablePartition {
		confReader.Partition = consumer.Partition
	} else {
		// GroupID 要从中读取消息的分区。可以分配Partition或GroupID中的一个，但不能同时分配两个
		confReader.GroupID = consumer.GroupID
	}
	r := kafka.NewReader(confReader)
	return r
}
