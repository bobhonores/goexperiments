package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"

	nt "github.com/bobhonores/goexperiments/internal/nats"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer nc.Drain()

	_, err := c.Subscribe("demo", func(t *nt.Temperature) {
		fmt.Printf("Received temperature: %v%v at %v\n", t.Degrees, t.Scale, time.Now())
	})

	if err != nil {
		log.Fatalf("Error when subscribing: %v", err)
		return
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
			time.Sleep(10 * time.Second)
			fmt.Println("Waiting ...")
		}
	}

	c.Close()
}
