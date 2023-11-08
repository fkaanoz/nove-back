package kafka

import (
	"fmt"
	confkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type TestTpc struct {
	CommitFrequency int
}

func (t TestTpc) HandleMessage(msg *confkafka.Message) error {
	fmt.Println("test-topic", string(msg.Value))
	return nil
}

// write a committer.
