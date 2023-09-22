package rmq

import (
	"consumer/utils"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

var OpenChannelsMap map[utils.MatchCriteria]*amqp.Channel
var ResultChannel *amqp.Channel
var Conn *amqp.Connection

func Init() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		msg := fmt.Sprintf("[Init] Error dialing TCP connection | err: %v", err)
		log.Println(msg)
		panic(err)
	}

	OpenChannelsMap = make(map[utils.MatchCriteria]*amqp.Channel, 4)
	Conn = connectRabbitMQ

	// Construct requests MQ
	for _, channelType := range utils.MatchCriterias {
		channelRabbitMQ, err := connectRabbitMQ.Channel()
		if err != nil {
			msg := fmt.Sprintf("[Init] Error creating unique criteria channel | err: %v", err)
			log.Println(msg)
			panic(err)
		}

		OpenChannelsMap[channelType] = channelRabbitMQ

		_, err = channelRabbitMQ.QueueDeclare(
			string(channelType), // queue name
			false,               // durable
			false,               // auto delete
			false,               // exclusive
			false,               // no wait
			nil,                 // arguments
		)
		if err != nil {
			msg := fmt.Sprintf("[Init] Error declaring criteria MQ | err: %v", err)
			log.Println(msg)
			panic(err)
		}
	}

	// Constructs result MQ
	mq, err := connectRabbitMQ.Channel()
	if err != nil {
		msg := fmt.Sprintf("[Init] Error creating unique results channel | err: %v", err)
		log.Println(msg)
		panic(err)
	}
	ResultChannel = mq
	_, err = ResultChannel.QueueDeclare(
		"results", // queue name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		msg := fmt.Sprintf("[Init] Error declaring results MQ | err: %v", err)
		log.Println(msg)
		panic(err)
	}
}

func Reset() {
	var err error
	for _, openChannel := range OpenChannelsMap {
		err = openChannel.Close()
		if err != nil {
			msg := fmt.Sprintf("[Reset] Error closing criteria channels | err: %v", err)
			log.Println(msg)
			panic(err)
		}
	}

	err = Conn.Close()
	if err != nil {
		msg := fmt.Sprintf("[Reset] Error closing TCP connection | err: %v", err)
		log.Println(msg)
		panic(err)
	}

	err = ResultChannel.Close()
	if err != nil {
		msg := fmt.Sprintf("[Reset] Error closing result channel | err: %v", err)
		log.Println(msg)
		panic(err)
	}
}
