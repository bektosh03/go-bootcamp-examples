package consumer

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
)

const topic = "registrations"

type StudentMarkedEvent struct {
	Mark      int32  `json:"mark"`
	StudentID string `json:"student_id"`
	JournalID string `json:"journal_id"`
}

type Consumer interface {
	Events() <-chan StudentMarkedEvent
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

func (c KafkaConsumer) Events() <-chan StudentMarkedEvent {
	ch := make(chan StudentMarkedEvent)

	go func() {
		defer close(ch)
		for msg := range c.consumer.Messages() {
			var event StudentMarkedEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Println("[ERROR] could not unmarshal:", err)
				continue
			}
			ch <- event
		}
	}()

	return ch
}
