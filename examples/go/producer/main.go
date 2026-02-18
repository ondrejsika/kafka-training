package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FlagBrokerAddr string
var FlagTopic string

var Cmd = &cobra.Command{
	Use:   "producer",
	Short: "Example Kafka producer",
	Run: func(cmd *cobra.Command, args []string) {
		brokerAddr := viper.GetString("broker_addr")
		if brokerAddr == "" {
			fmt.Fprintln(os.Stderr, "error: broker-addr is required (--broker-addr flag, BROKER_ADDR env var, or .env file)")
			os.Exit(1)
		}
		ctx := context.Background()
		produce(ctx, brokerAddr, FlagTopic)
	},
}

func init() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	Cmd.Flags().StringVarP(
		&FlagBrokerAddr,
		"broker-addr",
		"b",
		"",
		"Broker address (env: BROKER_ADDR)",
	)
	_ = viper.BindPFlag("broker_addr", Cmd.Flags().Lookup("broker-addr"))
	Cmd.Flags().StringVarP(
		&FlagTopic,
		"topic",
		"t",
		"",
		"Topic name",
	)
	Cmd.MarkFlagRequired("topic")
}

func main() {
	Cmd.Execute()
}

func produce(
	ctx context.Context,
	brokerAddr string,
	topic string,
) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddr},
		Topic:   topic,
	})

	i := 0
	var key string
	var msg string

	for {
		key = strconv.Itoa(i)
		msg = "hello_" + strconv.Itoa(i)
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(key),
			Value: []byte(msg),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}
		fmt.Printf("produce: topic=%s key=%s msg=%s\n", topic, key, msg)
		i++
		time.Sleep(time.Second)
	}
}
