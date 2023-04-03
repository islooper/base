package mq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type ConsumerClientConf struct {
	GroupName string
	Address   string
	Topic     string
}

func NewConsumerClient(c ConsumerClientConf) (rocketmq.PushConsumer, error) {
	// 消息主动推送给消费者
	cli, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(c.GroupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{c.Address})),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset), // 选择消费时间(首次/当前/根据时间)
		consumer.WithConsumerModel(consumer.BroadCasting))              // 消费模式(集群消费:消费完其他人不能再读取/广播消费：所有人都能读)

	if err != nil {
		return nil, err
	}

	return cli, nil
}
