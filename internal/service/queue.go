package service

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf(fmt.Sprintf("%s: %s", msg, err))
	}
}
func NewQueueConn() *amqp.Connection {
	ru := os.Getenv("RABBITMQ_DEFAULT_USER")
	rp := os.Getenv("RABBITMQ_DEFAULT_PASS")
	url := fmt.Sprintf("amqp://%s:%s@localhost:5672/", ru, rp)
	con, err := amqp.Dial(url)

	failOnError(err, "Failed to connect to RabbitMQ")

	return con
}
