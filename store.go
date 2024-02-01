package gomb

import (
	"fmt"
	"net"
)

type Store struct {
	topics  map[string]*Topic
	clients []net.Conn // TODO: Find a better abstraction for client
}

func NewStore() *Store {
	return &Store{
		topics:  make(map[string]*Topic),
		clients: []net.Conn{},
	}
}

func validateName(name string) bool {
	// Implement name validation
	return true
}

func (s *Store) CreateTopic(name string) error {
	// TODO: Implement sync (through a mutex ?)
	_, ok := s.topics[name]
	if ok {
		return fmt.Errorf("A topic with name %s already exists", name)
	}
	s.topics[name] = NewTopic(name)
	return nil
}

func (s *Store) GetTopics() []*TopicListItem {
	topics := []*TopicListItem{}
	for _, topic := range s.topics {
		topics = append(topics, &TopicListItem{Name: topic.name, Messages: len(topic.messages), Subscribers: len(topic.subscribers)})
	}
	return topics
}
