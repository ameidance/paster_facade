package client

import (
    "context"
    "time"

    "github.com/ameidance/paster_facade/conf"
    "github.com/go-redis/redis/v8"
)

var (
    RedisClient *redis.Client
)

func InitRedis() {
    redisConf, err := conf.GetRedisConfig()
    if redisConf == nil || err != nil {
        panic(err)
    }

    RedisClient = redis.NewClient(&redis.Options{
        Addr:     redisConf.Addr,
        Password: redisConf.Password,
        DB:       redisConf.DB,
        PoolSize: redisConf.PoolSize,
    })

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if _, err = RedisClient.Ping(ctx).Result(); err != nil {
        panic(err)
    }
}
