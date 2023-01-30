package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	// write to kafka topic "notifications"
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":            "localhost:29092",
		"client.id":                    "go-producer" + time.Now().String(),
		"acks":                         "all",
		"retries":                      0,
		"linger.ms":                    1,
		"compression.type":             "snappy",
		"batch.num.messages":           1000,
		"queue.buffering.max.messages": 100000,
		"queue.buffering.max.ms":       1000,
		"message.send.max.retries":     3,
		"retry.backoff.ms":             5,
		"socket.keepalive.enable":      true,
		"socket.nagle.disable":         true,
		"socket.max.fails":             3,
		"broker.address.ttl":           1000,
		"broker.address.family":        "v4",
		"api.version.request":          true,
		"api.version.fallback.ms":      0,
		"security.protocol":            "plaintext",
		"ssl.key.location":             "",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	topic := "notifications"
	for {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte("Hello World" + strconv.Itoa(time.Now().Second())),
		}, nil)
		time.Sleep(1 * time.Second)
	}
}
