package main

import (
	"os"

	"github.com/streadway/amqp"
)

func main() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	println("Successfully connected to RabbitMQ instance")
}
