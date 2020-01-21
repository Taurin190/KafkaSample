package consumer

import (
	"flag"
	"fmt"
)

var (
	mongoServers = flag.String("mongoServers", "localhost:27017", "kafka")
)

type consumePreserver struct {
	consumer
}

func GetConsumerPreserver() Consumer {
	return &consumePreserver{}
}

func (c *consumePreserver) exec(consumedMessage ConsumedMessage) {
	fmt.Println(fmt.Sprintf("consumed message. message: %s, timestamp: %d", consumedMessage.Message, consumedMessage.Timestamp))
}
