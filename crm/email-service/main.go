package main

import (
	"github.com/Shopify/sarama"
	"log"
	"net"
)

func main() {
	cfg, err := Load()
	if err != nil {
		panic(err)
	}

	kafkaCfg := sarama.NewConfig()
	kafkaCfg.Consumer.Return.Errors = true
	client, err := sarama.NewClient(
		[]string{net.JoinHostPort(cfg.KafkaHost, cfg.KafkaPort)},
		kafkaCfg,
	)
	if err != nil {
		panic(err)
	}
	consumer, err := NewConsumer(client)
	if err != nil {
		panic(err)
	}

	defer client.Close()
	service := NewService(LoggingEmailSender{}, consumer)
	service.Run()
}

type LoggingEmailSender struct{}

func (s LoggingEmailSender) Send(email Email) error {
	log.Printf("Send an email to %s, with body: %s\n", email.To, email.Body)
	return nil
}
