package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const TOPIC = "example_simple"

func main() {
	brokerAddr := os.Getenv("BROKER_ADDR")
	if brokerAddr == "" {
		brokerAddr = "127.0.0.1:9092"
	}

	w := &kafka.Writer{
		Addr:                   kafka.TCP(brokerAddr),
		Topic:                  TOPIC,
		AllowAutoTopicCreation: true,
	}

	for i := 0; ; i++ {
		key := strconv.Itoa(i)
		msg := "simple_" + strconv.Itoa(i)
		err := w.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(key),
			Value: []byte(msg),
		})
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("produced: key=%s msg=%s\n", key, msg)
		time.Sleep(time.Second)
	}
}
