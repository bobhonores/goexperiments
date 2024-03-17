package main

import (
	"fmt"
	"math/rand"
	"os"

	kinter "github.com/bobhonores/goexperiments/internal/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path>\n", os.Args[0])
		os.Exit(1)
	}

	configFilePath := os.Args[1]
	config, err := kinter.Read(configFilePath)
	if err != nil {
		fmt.Printf("Failed to read config file: %s", err)
		os.Exit(1)
	}

	topic := "purchases"
	p, err := kafka.NewProducer(config)
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	users := []string{"rhonores", "lreyes", "sochoa", "rnaupari"}
	items := []string{"book", "gift card", "alarm clock", "t-shirts"}

	for n := 0; n < 10; n++ {
		key := users[rand.Intn(len(users))]
		data := items[rand.Intn(len(items))]
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(data),
		}, nil)
	}

	// Wait for all message to be delivered
	p.Flush(15 * 1000)
	p.Close()
}
