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
		if brokerAddr == "" {
			fmt.Fprintln(os.Stderr, "error: broker-addr is required (--broker-addr flag, BROKER_ADDR env var, or .env file)")
			os.Exit(1)
		}
		ctx := context.Background()
		consume(ctx, brokerAddr, FlagTopic, FlagGroupID)
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
	Cmd.Flags().StringVarP(
		&FlagGroupID,
		"group-id",
		"g",
		"",
		"Consumer group ID",
	)
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
