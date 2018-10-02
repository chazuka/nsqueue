package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

func produce(topic string, address string) {
	p := getProducer(address)
	for i := 0; i < 1000; i++ {
		if err := p.Publish(topic, []byte(fmt.Sprintf("id:%d", i))); err != nil {
			log.Printf("publishing: %d; err: %v", i, err)
		}
	}
	log.Println("publishing done")
}

func getProducer(address string) *nsq.Producer {
	config := nsq.NewConfig()
	config.MaxAttempts = 1
	config.MaxInFlight = 100

	producer, err := nsq.NewProducer(address, config)
	if err != nil {
		log.Fatal(err)
	}

	return producer
}
