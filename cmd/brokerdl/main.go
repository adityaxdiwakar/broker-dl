package main

import (
	"context"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	brokerdl "github.com/adityaxdiwakar/broker-dl"
)

var conf tomlConfig

func init() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatalf("error: could not parse configuration: %v\n", err)
	}
}

func main() {
	r := brokerdl.GetKafkaReader(conf.KafkaUrl)

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n",
			m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		r.CommitMessages(ctx, m)
	}
}
