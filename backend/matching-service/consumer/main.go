package main

import (
	"consumer/rmq"
	"consumer/utils"
	"consumer/worker"
)

func main() {
	rmq.Init()
	defer rmq.Reset()
	forever := make(chan bool)

	// Sets off worker goroutines to listen at each criteria channel
	for _, criteria := range utils.MatchCriterias {
		worker.SpinMQConsumer(criteria)
	}

	<-forever
}
