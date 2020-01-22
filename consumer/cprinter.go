package consumer

import (
	"fmt"

	"../config"
)

type consumePrinter struct {
	Config config.Config
}

func GetConsumerPrinter(c config.Config) *consumePrinter {
	return &consumePrinter{
		Config: c,
	}
}

func (c *consumePrinter) exec(consumedMessage ConsumedMessage) {
	fmt.Println(fmt.Sprintf("consumed message. message: %s, timestamp: %d", consumedMessage.Message, consumedMessage.Timestamp))
}

func (c *consumePrinter) Run() {
	consumer := GetConsumer(c.Config)
	consumer.Run(c.exec)
}
