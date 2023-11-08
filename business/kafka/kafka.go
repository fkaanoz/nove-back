package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// subscribed topics.
const (
	TestTopic    = "test-topic"
	AnotherTopic = "another-topic"
)

type MessageHandler interface {
	HandleMessage(msg *kafka.Message) error
}

var TopicMap = map[string]MessageHandler{
	AnotherTopic: AnotherTpc{
		CommitFrequency: 10,
	},
	TestTopic: TestTpc{
		CommitFrequency: 20,
	},
}
