package c

import (
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
)

func sendMessage(brokerAddr []string, config *sarama.Config, topic string, value sarama.Encoder) {
	producer, err := sarama.NewSyncProducer(brokerAddr, config)
	if err != nil {
		fmt.Println("33333", err)
		return
	}
	defer func() {
		if err = producer.Close(); err != nil {
			fmt.Println("1111", err)
			return
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: value,
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		fmt.Println("2222", err)
		return
	}
	fmt.Printf("partition:%d offset:%d\n", partition, offset)
}

func Test_Kafka(t *testing.T) {
	ips := []string{"**.**.76.120:9092"}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	sendMessage(ips, config, "test-topic", sarama.ByteEncoder("haha"))
}
