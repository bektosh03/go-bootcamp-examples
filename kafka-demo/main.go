package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	fmt.Println("I'm a producer")

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true

	client, err := sarama.NewClient([]string{"localhost:9092"}, cfg)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	broker, err := client.Controller()
	if err != nil {
		panic(err)
	}

	_, err = broker.CreateTopics(&sarama.CreateTopicsRequest{
		Version: 1,
		TopicDetails: map[string]*sarama.TopicDetail{
			"greetings": {
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
