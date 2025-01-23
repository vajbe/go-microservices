package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"go-microservices/products/db"
	"go-microservices/products/types"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader *kafka.Reader
}

// NewKafkaConsumer initializes a new Kafka consumer
func NewKafkaConsumer(brokers []string, topic, groupID string) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topic,
		GroupID:   groupID,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		Partition: 0,
	})
	return &KafkaConsumer{
		Reader: reader,
	}
}

// Consume listens for messages from the Kafka topic
func (kc *KafkaConsumer) Consume(ctx context.Context) {
	for {
		msg, err := kc.Reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error while reading message: %v\n", err)
			continue
		}
		log.Printf("In Product Service Message consumed: key=%s, value=%s\n", string(msg.Key), string(msg.Value))
		handleProductUpdates(msg.Value)
	}
}

func handleProductUpdates(msg []byte) {
	var products types.Order
	err := json.Unmarshal(msg, &products)
	if err != nil {
		fmt.Printf("\n err %s", err.Error())
	}

	for _, el := range products.Products {
		product, err := db.GetProductById(el.ProductID)
		if err != nil {
			fmt.Printf("\n error occurred while fetching product details: %s", err.Error())
		}
		fmt.Printf("\n%+v", product)
	}

}

// Close gracefully shuts down the consumer
func (kc *KafkaConsumer) Close() {
	if kc.Reader != nil {
		_ = kc.Reader.Close()
		log.Println("Kafka consumer closed.")
	}
}
