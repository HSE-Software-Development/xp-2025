package manager

import (
	"log"

	"encoding/json"
	"github.com/HSE-Software-Development/xp-2025/internal/utils"
	"github.com/IBM/sarama"
)

type KafkaSimple struct {
	admin    sarama.ClusterAdmin
	producer sarama.SyncProducer
	consumer sarama.Consumer
	topic    string
}


func New(brokers []string) (*KafkaSimple, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Producer.Return.Successes = true

	admin, err := sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		return nil, err
	}

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		admin.Close()
		return nil, err
	}

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		admin.Close()
		producer.Close()
		return nil, err
	}

	return &KafkaSimple{
		admin:    admin,
		producer: producer,
		consumer: consumer,
	}, nil
}

func (k *KafkaSimple) CreateTopic(topic string) error {
	return k.admin.CreateTopic(topic, &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false)
}

func (k *KafkaSimple) Subscribe(topic string, handler chan utils.Message) error {
	partitionConsumer, err := k.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}

	k.topic = topic

	go func() {
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				var message utils.Message
				if err := json.Unmarshal(msg.Value, &message); err != nil {
					log.Printf("Failed to unmarshal message: %v", err)
					continue
				}
				handler <- message
			case err := <-partitionConsumer.Errors():
				log.Printf("Kafka error: %v", err)
			}
		}
	}()

	return nil
}

func (k *KafkaSimple) Send(message utils.Message) error {
	if k.topic == "" {
		return nil
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, _, err = k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: k.topic,
		Value: sarama.ByteEncoder(jsonData),
	})

	return err
}

func (k *KafkaSimple) Close() {
	if k.admin != nil {
		k.admin.Close()
	}
	if k.producer != nil {
		k.producer.Close()
	}
	if k.consumer != nil {
		k.consumer.Close()
	}
}