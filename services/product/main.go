package main

import (
	"context"
	"fmt"
	"go-microservices/products/config"
	"go-microservices/products/db"
	"go-microservices/products/kafka"
	"go-microservices/products/middleware"
	"go-microservices/products/routes"
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

	// Initialize Table and dbs
	db.InitializeDb(SERVICE_CONFIG)

	//	Initialize Redis
	db.InitRedis(SERVICE_CONFIG)

	// Initialize Kafka (producer and consumer)
	kafkaCtx, kafkaCancel := context.WithCancel(context.Background())
	defer kafkaCancel()
	producer, consumer := InitKafka(kafkaCtx)
	kafka.SetKafkaManager(consumer, producer)

	// Initialize router
	// Initialize router
	router := mux.NewRouter()

	// Apply middlewares
	router.Use(middleware.LoggingMiddleware)

	// Register routes
	routes.RegisterProductRoutes(router)

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

	log.Printf("Product Service running on port %s", SERVICE_CONFIG.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func InitKafka(ctx context.Context) (*kafka.KafkaProducer, *kafka.KafkaConsumer) {
	// Kafka broker and topic configuration
	brokers := []string{"localhost:9092"}
	topic := "products"

	// Initialize producer
	producer := kafka.NewKafkaProducer(brokers, topic)

	// Initialize consumer
	consumer := kafka.NewKafkaConsumer(brokers, "orders", "product-group")

	// Start consuming in the background
	go consumer.Consume(ctx)
	fmt.Println("Connected to kafka")
	return producer, consumer
}
