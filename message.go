package gomb

import (
	"time"
)

const (
	MessageConnect    = "connect"
	MessageDisconnect = "disconnect"
	MessageData       = "data"
)

type Message struct {
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	Topic     string    `json:"topic"`
}
