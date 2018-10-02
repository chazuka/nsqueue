package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

func getConsumer(topic, channel string) *nsq.Consumer {
	config := nsq.NewConfig()
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}
	c.AddConcurrentHandlers(nsq.HandlerFunc(handleMessage), 2)
	return c
}

func handleMessage(message *nsq.Message) error {
	log.Printf("Message ID: %s", message.ID)
	log.Println(string(message.Body))
	return nil
}

func consume(topic, channel, address string) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	c := getConsumer(topic,channel)
	log.Printf("start subscribing to channel:%s; topic:%s", channel, topic)
	if err := c.ConnectToNSQD(address); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}

