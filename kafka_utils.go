package caliber

import (
	"time"

	"github.com/segmentio/kafka-go"
)

func GetKafkaWriter(brokerList []string, topic string, async bool) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:       brokerList,
		Topic:         topic,
		Balancer:      &kafka.LeastBytes{},
		BatchTimeout:  time.Millisecond * 5,
		QueueCapacity: 10000,
		Async:         async,
	})
}

func GetKafkaReader(brokerList []string, topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokerList,
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e2, // 1KB
		MaxBytes:  10e6, // 10MB
	})
}
