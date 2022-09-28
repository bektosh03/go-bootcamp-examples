package main

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"strings"
	"time"
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

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	for {
		_, _, err = producer.SendMessage(generateMessage())
		if err != nil {
			fmt.Println("failed to send message:", err)
		}
	}
}

func generateMessage() *sarama.ProducerMessage {
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your greeting: ")
	greeting, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	return &sarama.ProducerMessage{
		Topic:     "greetings",
		Key:       sarama.StringEncoder(name),
		Value:     sarama.StringEncoder(strings.TrimSpace(greeting)),
		Offset:    1,
		Partition: 0,
		Timestamp: time.Now(),
	}
}
