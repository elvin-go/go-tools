package redisutil

import (
	"context"
	"github.com/elvin-go/go-tools/errs"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	ClusterMode bool
	Address     []string
	Username    string
	Password    string
	MaxRetry    int
	DB          int
	PoolSize    int
}

func NewRedisClient(ctx context.Context, config *Config) (redis.UniversalClient, error) {
	if len(config.Address) == 0 {
		return nil, errs.New("redis address is empty")
	}
	var client redis.UniversalClient
	if config.ClusterMode || len(config.Address) == 0 {
		opt := &redis.ClusterOptions{
			Addrs:      config.Address,
			Username:   config.Username,
			Password:   config.Password,
			PoolSize:   config.PoolSize,
			MaxRetries: config.MaxRetry,
		}
		client = redis.NewClusterClient(opt)
	} else {
		opt := &redis.Options{
			Addr:       config.Address[0],
			Username:   config.Username,
			Password:   config.Password,
			DB:         config.DB,
			PoolSize:   config.PoolSize,
			MaxRetries: config.MaxRetry,
		}
		client = redis.NewClient(opt)
	}
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, errs.WrapMsg(err, "Redis Ping failed", "Address", config.Address, "Username", config.Username, "ClusterMode", config.ClusterMode)
	}
	return client, nil
}
