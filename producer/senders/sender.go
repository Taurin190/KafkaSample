package sender

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

type KafkaSender struct{}

var (
	kafkaServers = flag.String("kafkaServers", "localhost:32770", "kafka address")
)

// SendMessage 送信メッセージ
type SendMessage struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func GetKafkaSender() *KafkaSender {
	return &KafkaSender{}
}

func (sender *KafkaSender) Send(text string) {
	flag.Parse()
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

	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 3

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}
	timestamp := time.Now().UnixNano()

	send := &SendMessage{
		Message:   text,
		Timestamp: timestamp,
	}

	jsBytes, err := json.Marshal(send)
	if err != nil {
		panic(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: "test.A",
		Key:   sarama.StringEncoder(strconv.FormatInt(timestamp, 10)),
		Value: sarama.StringEncoder(string(jsBytes)),
	}

	producer.Input() <- msg

	select {
	case <-producer.Successes():
		fmt.Println(fmt.Sprintf("success send. message: %s, timestamp: %d", send.Message, send.Timestamp))
	case err := <-producer.Errors():
		fmt.Println(fmt.Sprintf("fail send. reason: %v", err.Msg))
	case <-ctx.Done():
		return
	}
}
