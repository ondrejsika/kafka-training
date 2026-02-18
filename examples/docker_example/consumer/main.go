package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

const TOPIC = "example_simple"
const GROUP_ID = "example_simple"

func main() {
	brokerAddr := os.Getenv("BROKER_ADDR")
	if brokerAddr == "" {
		brokerAddr = "127.0.0.1:9092"
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddr},
		Topic:   TOPIC,
		GroupID: GROUP_ID,
	})

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("consumed: key=%s msg=%s\n", msg.Key, msg.Value)
	}
}
