package main

import (
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FlagBrokerAddr string
var FlagTopic string
var FlagGroupID string

var Cmd = &cobra.Command{
	Use:   "consumer",
	Short: "Example Kafka consumer",
	Run: func(cmd *cobra.Command, args []string) {
		brokerAddr := viper.GetString("broker_addr")
		topic := viper.GetString("topic")
		groupID := viper.GetString("group_id")
		if brokerAddr == "" {
			fmt.Fprintln(os.Stderr, "error: broker-addr is required (--broker-addr flag, BROKER_ADDR env var, or .env file)")
			os.Exit(1)
		}
		if topic == "" {
			fmt.Fprintln(os.Stderr, "error: topic is required (--topic flag, TOPIC env var, or .env file)")
			os.Exit(1)
		}
		ctx := context.Background()
		consume(ctx, brokerAddr, topic, groupID)
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
		"Topic name (env: TOPIC)",
	)
	_ = viper.BindPFlag("topic", Cmd.Flags().Lookup("topic"))
	Cmd.Flags().StringVarP(
		&FlagGroupID,
		"group-id",
		"g",
		"",
		"Consumer group ID (env: GROUP_ID)",
	)
	_ = viper.BindPFlag("group_id", Cmd.Flags().Lookup("group-id"))
}

func main() {
	Cmd.Execute()
}

func consume(
	ctx context.Context,
	brokerAddr string,
	topic string,
	groupID string,
) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddr},
		Topic:   topic,
		GroupID: groupID,
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Printf("consume: topic=%s key=%s msg=%s\n", topic, msg.Key, msg.Value)
	}
}
