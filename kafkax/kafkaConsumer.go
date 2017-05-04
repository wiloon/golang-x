package kafkax

import (
	"github.com/Shopify/sarama"
	"log"
)

func KafkaConsumerSaramaX() {
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

	topic := "topic0"
	// How to decide partition, is it fixed value...?
	index := 84
	log.Println("consume partition.")
	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	for {
		log.Println("consume msg")
		msg := consumer.Messages()
		s := *<-msg

		log.Println("consume msg:", string(s.Value))
		index++
	}

}
