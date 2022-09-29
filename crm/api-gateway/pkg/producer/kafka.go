package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

const (
	topic = "registrations"
)

func NewKafkaProducer(client sarama.Client) (KafkaProducer, error) {
	if err := initTopics(client); err != nil {
		return KafkaProducer{}, fmt.Errorf("could not init topics: %w", err)
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return KafkaProducer{}, fmt.Errorf("could not create new sync producer from client: %w", err)
	}

	return KafkaProducer{
		producer: producer,
	}, nil
}

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func (p KafkaProducer) Produce(event RegisteredEvent) error {
	_, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(event.Email),
		Value:     event,
		Timestamp: time.Now(),
	})

	return err
}

func initTopics(client sarama.Client) error {
	topics, err := client.Topics()
	if err != nil {
		return err
	}

	for _, t := range topics {
		if t == topic {
			return nil
		}
	}

	broker, err := client.Controller()
	if err != nil {
		return err
	}

	_, err = broker.CreateTopics(&sarama.CreateTopicsRequest{
		Version: 1,
		TopicDetails: map[string]*sarama.TopicDetail{
			topic: {
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		},
	})

	return err
}
