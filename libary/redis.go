package libary

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	gRedis "github.com/go-redis/redis"
	"time"
)

type redisClient = gRedis.Options

type redisConf struct {
	RedisClient *redisClient
}

// RedisInit redis初始化
func RedisInit(tomlName string) (client *gRedis.Client) {
	redisConf := &redisConf{}
	if err := paladin.Get(tomlName).UnmarshalTOML(redisConf); err != nil {
		panic(err)
	}
	client = gRedis.NewClient(&gRedis.Options{
		Addr:         redisConf.RedisClient.Addr,
		Password:     redisConf.RedisClient.Password,
		DB:           redisConf.RedisClient.DB,
		MaxRetries:   redisConf.RedisClient.MaxRetries,
		ReadTimeout:  redisConf.RedisClient.ReadTimeout * time.Second,
		WriteTimeout: redisConf.RedisClient.WriteTimeout * time.Second,
		PoolSize:     redisConf.RedisClient.PoolSize,
		MinIdleConns: redisConf.RedisClient.MinIdleConns,
		IdleTimeout:  redisConf.RedisClient.IdleTimeout * time.Second,
	})
	return
}
