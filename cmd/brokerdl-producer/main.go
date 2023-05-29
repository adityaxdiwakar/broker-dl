package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/BurntSushi/toml"
	brokerdl "github.com/adityaxdiwakar/broker-dl"
	"github.com/nxadm/tail"
	"github.com/segmentio/kafka-go"
)

var conf tomlConfig

func init() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatalf("error: could not parse configuration: %v\n", err)
	}
}

func main() {
	writer := brokerdl.GetKafkaWriter(conf.KafkaUrl)
	ctx := context.Background()

	t, err := tail.TailFile("log.txt", tail.Config{
		Follow: true,
		Location: &tail.SeekInfo{
			Whence: io.SeekEnd,
		},
	})
	if err != nil {
		log.Fatalf("error: could not open file to tail from: %v\n", err)
	}

	fmt.Println("Starting producing...")
	for line := range t.Lines {
		writer.WriteMessages(ctx, kafka.Message{
			Value: []byte(line.Text),
		})
	}
}
