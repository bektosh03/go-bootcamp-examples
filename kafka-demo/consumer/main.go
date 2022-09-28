package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	fmt.Println("I'm a consumer")
	cfg := sarama.NewConfig()
	cfg.Consumer.Return.Errors = true

	client, err := sarama.NewClient([]string{"localhost:9092"}, sarama.NewConfig())
	if err != nil {
		panic(err)
	}

	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partConsumer, err := consumer.ConsumePartition("greetings", 0, 0)
	if err != nil {
		panic(err)
	}
	defer partConsumer.Close()

	for msg := range partConsumer.Messages() {
		fmt.Printf("%s: %s - %v\n", msg.Key, msg.Value, msg.Timestamp)
	}
}
