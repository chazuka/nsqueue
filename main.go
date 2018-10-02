package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	TaskPublish = "pub"
	TaskSubscribe = "sub"
)

var (
	task = flag.String("task", "", fmt.Sprintf("task to run %s/%s",TaskPublish, TaskSubscribe))
	topic = flag.String("topic", "", "topic to engage")
	channel = flag.String("channel", "", "channel to engage")
	address = flag.String("address", "", "network address to bound to")
)

func main() {
	flag.Parse()

	if len(*topic) == 0 {
		log.Fatal("--topic required")
	}

	if len(*address) == 0 {
		log.Fatal("--address required")
	}

	switch *task {
	case TaskPublish:
		produce(*topic, *address)
	case TaskSubscribe:
		if len(*channel) == 0 {
			log.Fatal("--channel required")
		}
		consume(*topic, *channel, *address)
	default:
		log.Printf("task %s is not supported", *task)
	}
}