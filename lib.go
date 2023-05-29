package brokerdl

import (
	"github.com/segmentio/kafka-go"
)

func GetKafkaReader(server string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{server},
		GroupID:   "downloader",
		Topic:     "media-transfers",
		Partition: 0,
		MinBytes:  0,
	})
}

func GetKafkaWriter(server string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{server},
		Topic:   "media-transfers",
	})
}
