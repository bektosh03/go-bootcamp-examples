package producer

import (
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

const (
	registrationsTopic = "registrations"
	marksTopic         = "marks"
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

func (p KafkaProducer) Produce(event interface{}) error {
	switch v := event.(type) {
	case RegisteredEvent:
		_, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
			Topic:     registrationsTopic,
			Key:       sarama.StringEncoder(v.Email),
			Value:     v,
			Timestamp: time.Now(),
		})

		return err
	case StudentMarkedEvent:
		_, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
			Topic:     marksTopic,
			Key:       sarama.StringEncoder(v.JournalID),
			Value:     v,
			Timestamp: time.Now(),
		})

		return err
	default:
		return errors.New("unrecognized event")
	}
}

func initTopics(client sarama.Client) error {
	topics, err := client.Topics()
	if err != nil {
		return err
	}

	for _, t := range topics {
		if t == registrationsTopic {
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
			registrationsTopic: {
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		},
	})

	return err
}
