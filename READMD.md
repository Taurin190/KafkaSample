# KafkaSample
kafka-dockerで立ち上げのでclientのPub/Subのサンプル。

以下を実装する。
- RestAPIでPostした内容でTopicを作成する
- TopicをSubscribeしてDBに保存する

## RestAPI
`go run producer/server.go`でAPIが立ち上がるようにした。

## ローカルのkafkaで立ち上げる
以下のドキュメントを参考にkakfaは立ち上げた。
http://masato.github.io/2015/04/28/kafka-in-docker-getting-started/