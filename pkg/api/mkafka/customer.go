package mkafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
)

func (k *Kafka) StartCustomer(topic string) {
	consumer, err := sarama.NewConsumer(k.servers, defaultConfig())
	if err != nil {
		panic(err)
		return
	}
	partition, err := consumer.ConsumePartition(topic, 0, -1)
	if err != nil {
		return
	}
	for s := range partition.Messages() {
		task, err := k.toTask(s.Value)
		if err != nil {
			fmt.Printf("invalid task [topic:%s]\n", topic)
			continue
		}
		GetMsgCenter().Handle(task)
	}
}

func (k *Kafka) toTask(value []byte) (*Task, error) {
	task := &Task{}
	err := json.Unmarshal(value, task)
	return task, err
}
