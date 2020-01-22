package sender

import (
	"encoding/json"
	"flag"
	"os"
	"strconv"
	"time"

	"../../config"
	"github.com/Shopify/sarama"
)

type KafkaSender struct {
	kafkaServers []string
	Config       config.Config
}

var (
	kafkaServers = flag.String("kafkaServers", "localhost:32776", "kafka address")
)

// SendMessage 送信メッセージ
type SendMessage struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func GetKafkaSender(c config.Config) *KafkaSender {
	return &KafkaSender{
		Config: c,
	}
}

func (sender *KafkaSender) Send(text, topic string) {
	if sender.Config.KafkaServers[0] == "" {
		os.Exit(1)
	}

	if topic == "" {
		topic = "topic.A"
	}

	brokers := sender.Config.KafkaServers
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
