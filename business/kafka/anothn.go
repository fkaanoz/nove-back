package kafka

import (
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"math/rand"
)

type AnotherTpc struct {
	CommitFrequency int
}

func (t AnotherTpc) HandleMessage(msg *kafka.Message) error {
	fmt.Println("another-topic (50% err) ", string(msg.Value))

	if rand.Intn(10) > 5 {
		fmt.Println("ERROR")
		return errors.New("handle err")
	}

	return nil
}

// write a committer.
