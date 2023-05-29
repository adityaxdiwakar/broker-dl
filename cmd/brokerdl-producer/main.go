package main

import (
	"context"
	"encoding/json"
	"os"

	brokerdl "github.com/adityaxdiwakar/broker-dl"
	"github.com/segmentio/kafka-go"
)

type DownloadNotification struct {
	Hash     string `json:"info_hash"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

func main() {
	notification := DownloadNotification{
		Hash:     os.Args[2],
		Name:     os.Args[3],
		Location: os.Args[4],
	}

	b, err := json.Marshal(notification)
	if err != nil {
		os.Exit(1)
	}

	writer := brokerdl.GetKafkaWriter(os.Args[1])
	ctx := context.Background()

	writer.WriteMessages(ctx, kafka.Message{
		Value: []byte(string(b)),
	})
}
