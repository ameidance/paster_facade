package manager

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/ameidance/paster_facade/client"
	"github.com/bytedance/gopkg/util/logger"
)

const (
	_EXPIRE_TIME = time.Second
)

func IsOverFrequencyLimit(ctx context.Context, ip string) (overLimit bool) {
	_, err := client.RedisClient.Get(ctx, ip).Result()
	if err != nil && err != redis.Nil {
		logger.CtxErrorf(ctx, "[IsOverFrequencyLimit] redis get failed. err:%v", err)
		return
	}
	// if exists
	if err != redis.Nil {
		logger.CtxWarnf(ctx, "[IsOverFrequencyLimit] frequency over limit. ip:%v", ip)
		overLimit = true
	}

	_, err = client.RedisClient.SetEX(ctx, ip, "", _EXPIRE_TIME).Result()
	if err != nil {
		logger.CtxErrorf(ctx, "[IsOverFrequencyLimit] redis set failed. err:%v", err)
		return
	}

	return
}
