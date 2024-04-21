package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"

	nt "github.com/bobhonores/goexperiments/internal/nats"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer nc.Drain()

	d := &nt.Temperature{Degrees: 42, Scale: "C"}
	c.Publish("demo", d)
	fmt.Printf("Publishing temperature at %v\n", time.Now())

	c.Close()
}
