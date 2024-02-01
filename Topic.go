package gomb

import (
	"net"
)

type Topic struct {
	name        string
	messages    []*Message
	subscribers []net.Conn // TODO: find a better abstraction for clients
}

type TopicListItem struct {
	Name        string `json:"name"`
	Messages    int    `json:"messages"`
	Subscribers int    `json:"subscribers"`
}

func NewTopic(name string) *Topic {
	return &Topic{
		name:        name,
		messages:    []*Message{},
		subscribers: []net.Conn{},
	}
}
