package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"../config"
	"github.com/Shopify/sarama"
)

/*
Consumer interface: Consumerで受け取ったメッセージを処理する部分のみ異なるのでExec()内で処理する
*/
type Consumer interface {
	Run(exec func(consumedMessage ConsumedMessage), processName, consumedTopic string)
}

type consumer struct {
	Config config.Config
}

func GetConsumer(c config.Config) Consumer {
	return &consumer{
		Config: c,
	}
}

// ConsumedMessage 受信メッセージ
type ConsumedMessage struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func (c *consumer) Run(exec func(consumedMessage ConsumedMessage), processName, consumedTopic string) {
	if c.Config.KafkaServers[0] == "" {
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	brokers := c.Config.KafkaServers
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

	partition, err := consumer.ConsumePartition(consumedTopic, 0, sarama.OffsetNewest)
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
				exec(consumed)
			case <-ctx.Done():
				break CONSUMER_FOR
			}
		}
	}()

	fmt.Println(processName + " start.")

	<-signals

	fmt.Println(processName + " stop.")
}
