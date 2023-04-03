package mq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type ProducerClientConf struct {
	NameSrvAddr string
	Topic       string
	GroupName   string
	Retry       int32
}

func NewProducerClient(c ProducerClientConf) (rocketmq.Producer, error) {
	addr, err := primitive.NewNamesrvAddr(c.NameSrvAddr)
	if err != nil {
		return nil, err
	}

	p, err := rocketmq.NewProducer(
		producer.WithGroupName(c.GroupName),
		producer.WithNameServer(addr),
		producer.WithCreateTopicKey(c.Topic),
		producer.WithRetry(int(c.Retry)))
	if err != nil {
		panic(err)
	}

	return p, nil
}
