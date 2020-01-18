package main

import "flag"

var (
	kafkaServers = flag.String("kafkaServers", "localhost:32770", "kafka address")
)

// ConsumedMessage 受信メッセージ
type ConsumedMessage struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
