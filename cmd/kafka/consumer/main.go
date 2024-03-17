package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	config.SetKey("go.application.rebalance.enable", true)

	c, err := kafka.NewConsumer(config)
	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	topic := "purchases"
	_ = c.SubscribeTopics([]string{topic}, nil)

	tp := kafka.TopicPartition{
		Topic:     &topic,
		Partition: 0,
		Offset:    kafka.Offset(0),
	}
	if err = c.Assign([]kafka.TopicPartition{tp}); err != nil {
		fmt.Printf("Error when assigning %v\n", err)
		os.Exit(1)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				continue
			}

			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s at offset %v\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value), ev.TopicPartition.Offset)
		}
	}

	c.Close()
}
