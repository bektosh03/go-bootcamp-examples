package main

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
)

const topic = "registrations"

type Consumer interface {
	Events() <-chan RegisteredEvent
}

func NewConsumer(client sarama.Client) (Consumer, error) {
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return nil, err
	}

	partConsumer, err := consumer.ConsumePartition(topic, 0, 0)
	if err != nil {
		return nil, err
	}

	return KafkaConsumer{consumer: partConsumer}, nil
}

type KafkaConsumer struct {
	consumer sarama.PartitionConsumer
}

func (c KafkaConsumer) Events() <-chan RegisteredEvent {
	ch := make(chan RegisteredEvent)

	go func() {
		defer close(ch)
		for msg := range c.consumer.Messages() {
			var event RegisteredEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Println("[ERROR] could not unmarshal:", err)
				continue
			}
			ch <- event
		}
	}()

	return ch
}
