package consumer

import (
	"fmt"
	"log"
	"time"

	"../config"
	"github.com/globalsign/mgo/bson"
	"gopkg.in/mgo.v2"
)

type Message struct {
	ID        bson.ObjectId `bson:"_id"`
	Message   string        `bson:"message"`
	Timestamp string        `bson:"timestamp"`
}

type consumePreserver struct {
	Config config.Config
}

func GetConsumerPreserver(c config.Config) *consumePreserver {
	return &consumePreserver{
		Config: c,
	}
}

func (c *consumePreserver) exec(consumedMessage ConsumedMessage) {
	mongoInfo := &mgo.DialInfo{
		Addrs:    []string{c.Config.MongoServer[1]},
		Timeout:  20 * time.Second,
		Database: c.Config.MongoServer[2],
		Username: "kafka",
		Password: "kafka",
		Source:   "kafka",
	}

	session, err := mgo.DialWithInfo(mongoInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	db := session.DB(c.Config.MongoServer[2])

	mes := &Message{
		ID:        bson.NewObjectId(),
		Message:   consumedMessage.Message,
		Timestamp: fmt.Sprintf("%d", consumedMessage.Timestamp),
	}
	col := db.C("message")
	if err := col.Insert(mes); err != nil {
		log.Fatalln(err)
	}
}

func (c *consumePreserver) Run() {
	consumer := GetConsumer(c.Config)
	consumer.Run(c.exec)
}
