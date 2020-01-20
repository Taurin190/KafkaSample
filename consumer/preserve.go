package consumer

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	mongoServers = flag.String("mongoServers", "localhost:27017", "kafka")
)

func preserve(kafkaServers *string) {
	// flag.Parse()

	if *kafkaServers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	brokers := strings.Split(*kafkaServers, ",")
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
				fmt.Println(fmt.Sprintf("consumed message. message: %s, timestamp: %d", consumed.Message, consumed.Timestamp))
			case <-ctx.Done():
				break CONSUMER_FOR
			}
		}
	}()

	fmt.Println("go-kafka-example start.")

	<-signals

	fmt.Println("go-kafka-example stop.")
}
