package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	// read from kafka topic "notifications"
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost:29092",
		"group.id":           "go-consumer",
		"session.timeout.ms": 6000,
		"enable.auto.commit": false,
	})
	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"notifications"}, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on \n%s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
