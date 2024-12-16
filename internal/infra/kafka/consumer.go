package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

// NewConsumer cria uma nova instância de Consumer
func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap, // Desreferenciamos o ponteiro para armazenar o valor na struct
		Topics:    topics,
	}
}
