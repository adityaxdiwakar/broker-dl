package brokerdl

import "github.com/segmentio/kafka-go"

func GetKafkaReader(server string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{server},
		GroupID:   "downloader",
		Topic:     "media-transfers",
		Partition: 0,
		MinBytes:  0,
	})
}
