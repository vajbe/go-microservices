package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

// NewKafkaProducer initializes a new Kafka producer
func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
		Async:        false,
	}

	return &KafkaProducer{
		Writer: writer,
	}
}

// Publish sends a message to the Kafka topic
func (kp *KafkaProducer) Publish(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := kp.Writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	})
	if err != nil {
		return err
	}

	log.Printf("Message sent to Kafka topic: key=%s, value=%s\n", key, value)
	return nil
}

// Close gracefully shuts down the producer
func (kp *KafkaProducer) Close() {
	if kp.Writer != nil {
		_ = kp.Writer.Close()
		log.Println("Kafka producer closed.")
	}
}
