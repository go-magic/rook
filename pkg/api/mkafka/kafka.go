package mkafka

import "sync"

var (
	kafka     *Kafka
	onceKafka sync.Once
)

type Kafka struct {
	servers []string
}

func GetKafkaInstance() *Kafka {
	onceKafka.Do(func() {
		kafka = &Kafka{}
	})
	return kafka
}

func (k *Kafka) SetServers(servers []string) {
	kafka.servers = servers
}
