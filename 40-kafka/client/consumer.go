package client

import (
	"context"
	"log"
	"time"

	"github.com/IBM/sarama"
)

// sarama consumer group
type ConsumerGroup struct {
	// sarama consumer group

	config *config
}

type config struct {
	AutoTopicCreation bool     `yaml:"auto-topic-creation"`
	Topic             string   `yaml:"topic"`
	ConsumerGroup     string   `yaml:"consumer-group"`
	BootstrapServers  []string `yaml:"bootstrap-servers" envconfig:"KAFKA_BOOTSTRAP_SERVERS"`
	Username          string   `yaml:"sasl.username" envconfig:"KAFKA_USERNAME"`
	Password          string   `yaml:"sasl.password" envconfig:"KAFKA_PASSWORD"`
}

func NewConsumerGroup() *ConsumerGroup {
	configs := sarama.NewConfig()
	configs.Version = sarama.V2_0_0_0 // specify appropriate Kafka version
	configs.Consumer.Offsets.AutoCommit.Enable = true
	configs.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, configs)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	consumer := ConsumerGroup{}
	ctx := context.Background()

	for {
		err = consumerGroup.Consume(ctx, []string{topic}, &consumer)
		if err != nil {
			log.Panicf("Error from consumer: %v", err)
		}
	}

	return &ConsumerGroup{}
}
