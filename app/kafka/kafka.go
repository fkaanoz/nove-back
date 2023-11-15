package kafka

import (
	rkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
	"shtil/business/kafka"
)

type CustomConsumer struct {
	*rkafka.Consumer
	Logger *zap.SugaredLogger
	Topics []string
}

func NewKafkaConsumer(config *rkafka.ConfigMap, log *zap.SugaredLogger, subscribedTopics []string) (*CustomConsumer, error) {
	c, err := rkafka.NewConsumer(config)

	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics(subscribedTopics, nil)
	if err != nil {
		return nil, err
	}

	return &CustomConsumer{
		Consumer: c,
		Logger:   log,
		Topics:   subscribedTopics,
	}, nil

}

func (c *CustomConsumer) Run() {

	for {
		ev := c.Consumer.Poll(1000)

		switch e := ev.(type) {
		case *rkafka.Message:
			switch *e.TopicPartition.Topic {
			case kafka.TestTopic:
				err := kafka.TopicMap[kafka.TestTopic].HandleMessage(e)
				if err == nil {
					c.Commit()
				}

			case kafka.AnotherTopic:
				err := kafka.TopicMap[kafka.AnotherTopic].HandleMessage(e)
				if err == nil {
					c.Commit()
				}
			}
		case rkafka.PartitionEOF:
			c.Logger.Errorw("KAFKA", "Message_Read_Error", e.Error)
		case rkafka.Error:
			c.Logger.Errorw("KAFKA", "Kafka Error", e.Error())
		default:
		}
	}
}
