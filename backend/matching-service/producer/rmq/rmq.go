package rmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"producer/utils"
)

var OpenChannelsMap map[utils.MatchCriteria]*amqp.Channel
var LengthChannelsMap map[utils.MatchCriteria]*amqp.Channel
var ResultChannel *amqp.Channel
var CancelChannel *amqp.Channel
var InQueueChannel *amqp.Channel
var Conn *amqp.Connection

const CancelQueueName = "cancelQueue"
const InQueueName = "inQueueQueue"

const CancelExchange = "cancels"
const InQueueExchange = "inqueue"

// Init Initialises all required connection and channels based on the MatchCriteria
// defined within the utils folder
func Init() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		msg := fmt.Sprintf("[Init] Error dialing TCP connection | err: %v", err)
		log.Println(msg)
		panic(err)
	}

	Conn = connectRabbitMQ

	OpenChannelsMap = make(map[utils.MatchCriteria]*amqp.Channel, 4)

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
			msg := fmt.Sprintf("[Init] Error declaring criteria queue instance | err: %v", err)
			log.Println(msg)
			panic(err)
		}
	}

	// Constructs result MQ
	mq, err := connectRabbitMQ.Channel()
	if err != nil {
		msg := fmt.Sprintf("[Init] Error creating unique result channel | err: %v", err)
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
		msg := fmt.Sprintf("[Init] Error declaring result queue instance | err: %v", err)
		log.Println(msg)
		panic(err)
	}

	LengthChannelsMap = make(map[utils.MatchCriteria]*amqp.Channel, 4)

	// Constructs length MQ
	for _, channelType := range utils.MatchCriterias {
		exchangeName := "length" + string(channelType)
		lengthChannelMQ, err := connectRabbitMQ.Channel()
		if err != nil {
			msg := fmt.Sprintf("[Init] Error creating unique criteria MQ length channel | err: %v", err)
			log.Println(msg)
			panic(err)
		}

		LengthChannelsMap[channelType] = lengthChannelMQ

		// Declare length exchange
		err = lengthChannelMQ.ExchangeDeclare(
			exchangeName,
			amqp.ExchangeFanout, // type
			true,                // durable
			false,               // auto-deleted
			false,               // internal
			false,               // no-wait
			nil,                 // arguments
		)
		if err != nil {
			msg := fmt.Sprintf("[Init] Error declaring length exchange for %s | err: %v", err, string(channelType))
			log.Println(msg)
			panic(err)
		}
	}

	// Constructs cancel MQ
	cancelMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		msg := fmt.Sprintf("[Init] Error creating unique cancel channel | err: %v", err)
		log.Println(msg)
		panic(err)
	}
	CancelChannel = cancelMQ
	// Declare cancel exchange
	err = CancelChannel.ExchangeDeclare(
		CancelExchange,      // name
		amqp.ExchangeFanout, // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		msg := fmt.Sprintf("[Init] Error declaring exchange | err: %v", err)
		log.Println(msg)
		panic(err)
	}
	declaredCancelQueue, err := CancelChannel.QueueDeclare(
		CancelQueueName, // name
		false,           // durable
		false,           // delete when unused
		true,            // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		msg := fmt.Sprintf("[Init] Error declaring cancel queue | err: %v", err)
		log.Println(msg)
		panic(err)
	}

	err = CancelChannel.QueueBind(
		declaredCancelQueue.Name, // queue name
		"",                       // routing key
		CancelExchange,           // exchange
		false,
		nil,
	)
	if err != nil {
		msg := fmt.Sprintf("[Init] Error binding cancel queue to exchange | err: %v", err)
		log.Println(msg)
		panic(err)
	}

	// Constructs InQueue MQ
	inQueueMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		msg := fmt.Sprintf("[Init] Error creating unique inQueue channel | err: %v", err)
		log.Println(msg)
		panic(err)
	}
	InQueueChannel = inQueueMQ
	// Declare cancel exchange
	err = InQueueChannel.ExchangeDeclare(
		InQueueExchange,     // name
		amqp.ExchangeFanout, // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		msg := fmt.Sprintf("[Init] Error declaring inQueue exchange | err: %v", err)
		log.Println(msg)
		panic(err)
	}
	declaredInQueueQueue, err := InQueueChannel.QueueDeclare(
		InQueueName, // name
		false,       // durable
		false,       // delete when unused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		msg := fmt.Sprintf("[Init] Error declaring inQueue queue | err: %v", err)
		log.Println(msg)
		panic(err)
	}

	err = InQueueChannel.QueueBind(
		declaredInQueueQueue.Name, // queue name
		"",                        // routing key
		InQueueExchange,           // exchange
		false,
		nil,
	)
	if err != nil {
		msg := fmt.Sprintf("[Init] Error binding inQueue queue to exchange | err: %v", err)
		log.Println(msg)
		panic(err)
	}
}

// Reset Closes all connections and channels to prevent leaks
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

	for _, lengthChannel := range LengthChannelsMap {
		err = lengthChannel.Close()
		if err != nil {
			msg := fmt.Sprintf("[Reset] Error closing length channels | err: %v", err)
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

	err = CancelChannel.Close()
	if err != nil {
		msg := fmt.Sprintf("[Reset] Error closing cancel channel | err: %v", err)
		log.Println(msg)
		panic(err)
	}

	err = InQueueChannel.Close()
	if err != nil {
		msg := fmt.Sprintf("[Reset] Error closing inQueue channel | err: %v", err)
		log.Println(msg)
		panic(err)
	}
}
