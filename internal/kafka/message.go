package kafka

import "time"

type Message struct {
	ID        int
	Content   string
	Timestamp time.Time
}
