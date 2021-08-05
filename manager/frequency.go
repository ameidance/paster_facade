package manager

import (
	"context"
	"time"

	"github.com/ameidance/paster_facade/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis/v8"
)

const (
	_EXPIRE_TIME = time.Second
)

func IsOverFrequencyLimit(ctx context.Context, ip string) (overLimit bool) {
	_, err := client.RedisClient.Get(ctx, ip).Result()
	if err != nil && err != redis.Nil {
		klog.Errorf("[IsOverFrequencyLimit] redis get failed. err:%v", err)
		return
	}
	// if exists
	if err != redis.Nil {
		klog.Warnf("[IsOverFrequencyLimit] frequency over limit. ip:%v", ip)
		overLimit = true
	}

	_, err = client.RedisClient.SetEX(ctx, ip, "", _EXPIRE_TIME).Result()
	if err != nil {
		klog.Errorf("[IsOverFrequencyLimit] redis set failed. err:%v", err)
		return
	}

	return
}
