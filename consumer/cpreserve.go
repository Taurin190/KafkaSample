package consumer

import (
	"fmt"

	"../config"
)

type consumePreserver struct {
	Config config.Config
}

func GetConsumerPreserver(c config.Config) *consumePreserver {
	return &consumePreserver{
		Config: c,
	}
}

func (c *consumePreserver) exec(consumedMessage ConsumedMessage) {
	fmt.Println(fmt.Sprintf("consumed message. message: %s, timestamp: %d", consumedMessage.Message, consumedMessage.Timestamp))
}

func (c *consumePreserver) Run() {
	consumer := GetConsumer(c.Config)
	consumer.Run(c.exec)
}
