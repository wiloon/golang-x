package kafkax

import (
	"github.com/Shopify/sarama"
	"fmt"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Specify brokers address. This is default one
	brokers := []string{"localhost:9092"}

	// Create new consumer
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	topic := "test"
	// How to decide partition, is it fixed value...?
	index := 84
	for {
		consumer, err := master.ConsumePartition(topic, 0, index)
		if err != nil {
			panic(err)
		}

		msg := consumer.Messages()
		s := *<-msg

		fmt.Println("consume msg:", string(s.Value))
		index++
	}

}
