package sender

import (
	"encoding/json"
	"flag"
	"os"
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

func (sender *KafkaSender) Send(text, topic string) {
	flag.Parse()
	if *kafkaServers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if topic == "" {
		topic = "topic.A"
	}

	brokers := strings.Split(*kafkaServers, ",")
	config := sarama.NewConfig()

	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 3
	config.Producer.RequiredAcks = sarama.WaitForAll

	producer, err := sarama.NewSyncProducer(brokers, config)
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
		Topic: topic,
		Key:   sarama.StringEncoder(strconv.FormatInt(timestamp, 10)),
		Value: sarama.StringEncoder(string(jsBytes)),
	}

	producer.SendMessage(msg)
	producer.Close()
}
