package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

/*
Consumer interface: Consumerで受け取ったメッセージを処理する部分のみ異なるのでExec()内で処理する
*/
type Consumer interface {
	exec(consumedMessage ConsumedMessage)
	Run(kafkaServers []string)
}

type consumer struct {
}

func GetConsumer() Consumer {
	return &consumer{}
}

// ConsumedMessage 受信メッセージ
type ConsumedMessage struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func (c *consumer) Run(kafkaServers []string) {
	if kafkaServers[0] == "" {
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	brokers := kafkaServers
	config := sarama.NewConfig()

	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			panic(err)
		}
	}()

	partition, err := consumer.ConsumePartition("test.A", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	// コンシューマールーチン
	go func() {
	CONSUMER_FOR:
		for {
			select {
			case msg := <-partition.Messages():
				var consumed ConsumedMessage
				if err := json.Unmarshal(msg.Value, &consumed); err != nil {
					fmt.Println(err)
				}
				c.exec(consumed)
			case <-ctx.Done():
				break CONSUMER_FOR
			}
		}
	}()

	fmt.Println("go-kafka-example start.")

	<-signals

	fmt.Println("go-kafka-example stop.")
}

func (c *consumer) exec(consumedMessage ConsumedMessage) {
	fmt.Println(fmt.Sprintf("consumed message. message: %s, timestamp: %d", consumedMessage.Message, consumedMessage.Timestamp))
}
