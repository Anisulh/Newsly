package utils

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

type UserInteraction struct {
    UserID    uint
    ContentID uint
    Type      string // e.g., "like", "view", "bookmark"
}

func SendMessage(brokers []string, topic string, interaction UserInteraction) error {
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true

    // Create a new synchronous producer
    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        return err
    }
    defer producer.Close()

    // Serialize your UserInteraction struct to JSON or another format
    messageBytes, err := json.Marshal(interaction)
    if err != nil {
        return err
    }

    // Construct a message
    msg := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.ByteEncoder(messageBytes),
    }

    // Send the message
    partition, offset, err := producer.SendMessage(msg)
    if err != nil {
        return err
    }

    log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
    return nil
}
