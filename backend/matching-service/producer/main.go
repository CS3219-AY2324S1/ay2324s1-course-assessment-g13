package main

import (
	"os"

	"github.com/labstack/echo/v4"
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

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	_, err = channelRabbitMQ.QueueDeclare(
		"matchingService", // queue name
		true,              // durable
		false,             // auto delete
		false,             // exclusive
		false,             // no wait
		nil,               // arguments
	)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/send", func(c echo.Context) error {
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("hello"),
		}
		if err := channelRabbitMQ.Publish(
			"",                // exchange
			"MatchingService", // queue name
			false,             // mandatory
			false,             // immediate
			message,           // message to publish
		); err != nil {
			return err
		}

		return nil
	})
	e.Start(":8080")
}
