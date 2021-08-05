package client

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
)

var (
	RedisClient *redis.Client
)

func InitRedis() {
	redisConf, err := getRedisConfig()
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

const (
	_REDIS_CONF_PATH = "conf/redis.yml"
)

type _RedisConf struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"pool_size"`
}

func getRedisConfig() (*_RedisConf, error) {
	conf := new(_RedisConf)
	file, err := ioutil.ReadFile(_REDIS_CONF_PATH)
	if err != nil {
		logger.Errorf("[getRedisConfig] open file failed. err:%v", err)
		return nil, err
	}
	if err = yaml.Unmarshal(file, conf); err != nil {
		logger.Errorf("[getRedisConfig] unmarshal file failed. err:%v", err)
		return nil, err
	}
	return conf, nil
}
