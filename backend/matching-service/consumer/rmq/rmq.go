package rmq

import (
	"consumer/models"
	"consumer/utils"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"io"
	"log"
	"net/http"
	"os"
)

var OpenChannelsMap map[utils.MatchCriteria]*amqp.Channel
var SyncChannelsMap map[utils.MatchCriteria]*amqp.Channel
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
	SyncChannelsMap = make(map[utils.MatchCriteria]*amqp.Channel, 4)
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

		// Constructs sync MQ for distributed workers
		syncChannelMQ, err := connectRabbitMQ.Channel()
		if err != nil {
			msg := fmt.Sprintf("[Init] Error creating unique sync channel | err: %v", err)
			log.Println(msg)
			panic(err)
		}
		SyncChannelsMap[channelType] = syncChannelMQ
		_, err = syncChannelMQ.QueueDeclare(
			string(channelType)+"sync", // queue name
			false,                      // durable
			false,                      // auto delete
			false,                      // exclusive
			false,                      // no wait
			nil,                        // arguments
		)
		if err != nil {
			msg := fmt.Sprintf("[Init] Error declaring sync MQ | err: %v", err)
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

	for _, syncChannel := range SyncChannelsMap {
		err = syncChannel.Close()
		if err != nil {
			msg := fmt.Sprintf("[Reset] Error closing sync channels | err: %v", err)
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

func GetQueueSize(queueName string) int64 {
	client := &http.Client{}
	url := os.Getenv("RMQ_QUEUE_URL") + queueName
	log.Printf("URL OF QUEUE: %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		msg := fmt.Sprintf("[GetQueueSize] Error creating new request | err: %v", err)
		log.Fatal(msg)
	}
	req.SetBasicAuth("guest", "guest")
	resp, err := client.Do(req)
	if err != nil {
		msg := fmt.Sprintf("[GetQueueSize] Error executing HTTP request | err: %v", err)
		log.Fatal(msg)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		msg := fmt.Sprintf("[GetQueueSize] Error reading response body | err: %v", err)
		log.Fatal(msg)
	}
	var queueSizeResponse models.MessageQueueLengthResponse
	err = json.Unmarshal(bodyText, &queueSizeResponse)
	if err != nil {
		msg := fmt.Sprintf("[GetQueueSize] Error unmarshaling content | err: %v", err)
		log.Fatal(msg)
	}
	return queueSizeResponse.MessageStats.Publish - queueSizeResponse.MessageStats.Ack
}
