package main

import (
	"consumer/rmq"
	"consumer/utils"
	"consumer/worker"
	"log"
)

func main() {
	rmq.Init()
	rmq.RunQueueListeners()
	defer rmq.Reset()
	forever := make(chan bool)

	// Sets off worker goroutines to listen at each criteria channel
	for _, criteria := range utils.MatchCriterias {
		worker.SpinMQConsumer(criteria)
	}

	log.Printf("Matching consumer service has started!\n")

	<-forever
}
