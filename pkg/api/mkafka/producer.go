package mkafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
)

func newProducerMessage(topic string, data interface{}) (*sarama.ProducerMessage, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(string(b))
	return msg, nil
}

func defaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	// 生产者配置
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	return config
}

func (k *Kafka) SendMessage(topic string, data interface{}) error {
	msg, err := newProducerMessage(topic, data)
	if err != nil {
		return err
	}
	// 连接kafka
	client, err := sarama.NewSyncProducer(k.servers, defaultConfig())
	if err != nil {
		return err
	}
	defer client.Close()
	// 发送消息
	partition, offset, err := client.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("partition:%v offset:%v\n", partition, offset)
	return nil
}
