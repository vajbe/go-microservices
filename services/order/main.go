package main

import (
	"context"
	"go-microservices/order/config"
	"go-microservices/order/db"
	"go-microservices/order/kafka"
	"go-microservices/order/middleware"
	"go-microservices/order/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

var (
	REDIS_CLIENT   *redis.Client
	SERVICE_CONFIG config.Config
)

func main() {
	// Load configuration
	SERVICE_CONFIG = config.Load()

	// Initialize Table and DBs
	db.InitializeDb(SERVICE_CONFIG)

	// Initialize Redis
	db.InitRedis(SERVICE_CONFIG)

	// Initialize Kafka (producer and consumer)
	kafkaCtx, kafkaCancel := context.WithCancel(context.Background())
	defer kafkaCancel()
	producer, consumer := InitKafka(kafkaCtx)
	kafka.SetKafkaManager(consumer, producer)

	// Initialize router
	router := mux.NewRouter()

	// Apply middlewares
	router.Use(middleware.LoggingMiddleware)

	// Register routes
	routes.RegisterOrderRoutes(router)

	// Start server
	server := &http.Server{
		Addr:    ":" + SERVICE_CONFIG.Port,
		Handler: router,
	}

	// Handle termination signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stopChan
		log.Println("Shutting down gracefully...")
		kafkaCancel() // Stop Kafka consumer
		consumer.Close()
		producer.Close()

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("HTTP server Shutdown: %v", err)
		}
	}()

	log.Printf("Order Service running on port %s", SERVICE_CONFIG.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func InitKafka(ctx context.Context) (*kafka.KafkaProducer, *kafka.KafkaConsumer) {
	// Kafka broker and topic configuration
	brokers := []string{SERVICE_CONFIG.Kafka_URL}
	topic := "orders"

	// Initialize producer
	producer := kafka.NewKafkaProducer(brokers, topic)

	// Initialize consumer
	consumer := kafka.NewKafkaConsumer(brokers, topic, "order-group")

	// Start consuming in the background
	go consumer.Consume(ctx)

	return producer, consumer
}
