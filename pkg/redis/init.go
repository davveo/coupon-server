package redis

import (
	"context"
	"fmt"
	"github.com/davveo/market-coupon/pkg/config"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
	once   sync.Once
)

// Init 初始化redis连接
func Init(conf *config.Redis) (*redis.Client, error) {
	once.Do(func() {
		addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: conf.Password,
			DB:       conf.DB,
		})
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return client, nil
}

func Client() *redis.Client {
	return client
}
