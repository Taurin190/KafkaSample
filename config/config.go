package config

type Config struct {
	KafkaServers []string
	MongoServer  []string
}

func GetConfig() *Config {
	return &Config{
		KafkaServers: []string{"kafkaServers", "localhost:32770", "kafka address"},
		MongoServer:  []string{"mongoServers", "localhost:27017", "kafka"},
	}
}
