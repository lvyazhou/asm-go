package repo

import (
	"asm_platform/domain/repository"
	kfk "asm_platform/infrastructure/pkg/database/kafka"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

type KafkaRepo struct {
	Topic   string `json:"topic"`
	GroupID string `json:"groupID"`
}

// NewKafkaRepo 实例化对象
func NewKafkaRepo(topic string, groupId string) *KafkaRepo {
	return &KafkaRepo{
		Topic:   topic,
		GroupID: groupId,
	}
}

var _ repository.KafkaRepository = &KafkaRepo{}

// WriteKafkaMessage 写入kafka消息
func (k KafkaRepo) WriteKafkaMessage(m map[string]string) error {
	// 初始化kafka product repo
	product := kfk.KfkProduct{Topic: k.Topic}
	w := product.NewKfkProduct()

	// 定义发送的消息信息
	var messages []kafka.Message
	if len(m) > 0 {
		for k, v := range m {
			var msg = kafka.Message{
				Key:   []byte(k),
				Value: []byte(v),
			}
			messages = append(messages, msg)
		}
	}
	var err error

	// attempt to create topic prior to publishing the message
	err = w.WriteMessages(context.Background(), messages...)
	if err != nil {
		slog.Errorf("unexpected error %v", err)
	}

	if err := w.Close(); err != nil {
		slog.Errorf("failed to close writer:", err)
	}
	return err
}

// WriteKafkaMessageList 写入kafka消息
func (k KafkaRepo) WriteKafkaMessageList(mapList []map[string]string) error {
	// 初始化kafka product repo
	product := kfk.KfkProduct{Topic: k.Topic}
	w := product.NewKfkProduct()

	// 定义发送的消息信息
	var messages []kafka.Message
	if len(mapList) > 0 {
		for _, value := range mapList {
			for k, v := range value {
				var msg = kafka.Message{
					Key:   []byte(k),
					Value: []byte(v),
				}
				messages = append(messages, msg)
			}
		}
	}
	slog.Infof("kafka message size %v", len(messages))
	var err error

	// attempt to create topic prior to publishing the message
	err = w.WriteMessages(context.Background(), messages...)
	if err != nil {
		slog.Errorf("unexpected error %v", err)
	}

	if err := w.Close(); err != nil {
		slog.Errorf("failed to close writer:", err)
	}
	return err
}

func (k KafkaRepo) ReadKafkaMessage() {
	// 获取partition
	partitions, err := kfk.Conn.ReadPartitions(k.Topic)
	if err != nil {
		slog.Errorf("kafka consumer partition count error %v", err.Error())
	}
	if len(partitions) > 0 {
		//var wg sync.WaitGroup
		for _, p := range partitions {
			fmt.Println("partition id is ", p.ID)
			consumer := kfk.KfkConsumer{
				Topic:           k.Topic,
				EnablePartition: true,
				Partition:       p.ID,
				GroupID:         k.GroupID,
			}
			r := consumer.NewKfkConsumer()
			fmt.Println(r.Lag())
			//r.SetOffset()
			//wg.Add(1)
			go func() {
				for {
					m, err := r.ReadMessage(context.Background())
					if err != nil {
						slog.Errorf("read kafka message error %v", err.Error())
					}
					fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
				}
				//wg.Done()
			}()
		}
		//wg.Wait()
	}

}

func (k KafkaRepo) ReadKafkaMessageByGroupId() {
	consumer := kfk.KfkConsumer{
		Topic:           k.Topic,
		EnablePartition: false,
		GroupID:         k.GroupID,
	}
	r := consumer.NewKfkConsumer()

	ctx := context.Background()
	var result []string
	var wg sync.WaitGroup
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		result = append(result, string(m.Value))
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
		fmt.Println("result size is ", len(result))
		// 满足 5 就扔出去
		if len(result)%5 == 0 {
			//fmt.Println("five results ", result)
			wg.Add(1)
			go handleMessage(result, &wg)
			// 初始化空
			result = []string{}
		}
	}
	wg.Wait()
	fmt.Println("never end")
}

func handleMessage(result []string, wg *sync.WaitGroup) {
	fmt.Println("handle message start ...")
	fmt.Println(result)
	//time.Sleep(5 * time.Second)
	wg.Done()
}
