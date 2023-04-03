package sharedconfig

import (
	"github.com/go-redis/redis/v8"
)

type RedisConf struct {
	Host         string
	Passwd       string
	DB           int `json:",default=0"`
	PoolSize     int `json:",default=10"` // 连接池
	MinIdleConns int `json:",default=10"` // 在启动阶段创建指定数量的Idle连接
}

// NewRedis 初始化redis cli
func NewRedis(conf RedisConf) *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:      "tcp",
		Addr:         conf.Host,
		Password:     conf.Passwd,
		DB:           conf.DB,
		PoolSize:     conf.PoolSize,
		MinIdleConns: conf.MinIdleConns,
	})
}
