package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"reflect"
)

type (
	RedisCliSetter interface {
		SetCli(cli *redis.Client)
		SetCtx(ctx context.Context)
	}

	RedisConnNewer interface {
		NewConn(ctx context.Context) *redis.Conn
	}
)

type RedisDao struct {
	cli *redis.Client
	ctx context.Context
}

// 初始化redis dao
func NewRedisDao(dao RedisCliSetter, ctx context.Context, cli *redis.Client) RedisCliSetter {
	dao.SetCtx(ctx)
	dao.SetCli(cli)
	return dao
}

// InjectRedisCli dao注入
// @param container: struct指针，必须要有redis.cli成员，至少一个RedisCliSetter成员
func InjectRedisCli(container interface{}) error {
	var cli *redis.Client
	var cliSetterList []RedisCliSetter

	mutReflector := reflect.ValueOf(container).Elem()
	fieldNum := mutReflector.NumField()

	for fieldIndex := 0; fieldIndex < fieldNum; fieldIndex++ {
		member := mutReflector.Field(fieldIndex)

		switch concrete := member.Interface().(type) {
		case *redis.Client:
			fmt.Println(concrete)
			cli = concrete
		case RedisCliSetter:
			cliSetterList = append(cliSetterList, concrete)
		default:
			continue
		}
	}

	if cli == nil {
		return fmt.Errorf("redis cli not exist")
	} else if len(cliSetterList) == 0 {
		return fmt.Errorf("RedisCliSetter not exist")
	}

	for _, cliSetter := range cliSetterList {
		cliSetter.SetCli(cli)
	}
	return nil
}

func (dao *RedisDao) SetCli(cli *redis.Client) {
	dao.cli = cli
}

func (dao *RedisDao) GetCli() *redis.Client {
	return dao.cli
}

func (dao *RedisDao) SetCtx(ctx context.Context) {
	dao.ctx = ctx
}

func (dao *RedisDao) GetCtx() context.Context {
	return dao.ctx
}

func (dao RedisDao) NewConn(ctx context.Context) *redis.Conn {
	return dao.cli.Conn(ctx)
}
