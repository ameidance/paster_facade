package client

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/ameidance/paster_facade/constant"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
)

var (
	RedisClient *redis.Client
)

func InitRedis() {
	conf, err := getRedisConfig()
	if conf == nil || err != nil {
		panic(err)
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
		PoolSize: conf.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err = RedisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}
}

type redisConf struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"pool_size"`
}

func getRedisConfig() (*redisConf, error) {
	conf := new(redisConf)
	file, err := ioutil.ReadFile(constant.REDIS_CONF_PATH)
	if err != nil {
		klog.Errorf("[getRedisConfig] open file failed. err:%v", err)
		return nil, err
	}
	if err = yaml.Unmarshal(file, conf); err != nil {
		klog.Errorf("[getRedisConfig] unmarshal file failed. err:%v", err)
		return nil, err
	}
	return conf, nil
}
