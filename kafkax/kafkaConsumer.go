package kafkax

import (
	"github.com/bsm/sarama-cluster"
	"fmt"
)

func KafkaConsumer() {
	brokers := []string{"localhost:9092"}
	topics := []string{"topic0"}
	consumer, err := cluster.NewConsumer(brokers, "", topics, nil)

	if err != nil {
		fmt.Println("err")
	}

	msg := consumer.Messages()
	// consumer.
	// fmt.Println("msg:", msg)
	fmt.Println("msg:", <-msg)
}
