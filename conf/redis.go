package conf

import (
	"github.com/bytedance/gopkg/util/logger"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const (
	_REDIS_CONF_PATH = "conf/redis.yml"
)

type RedisConf struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"pool_size"`
}

func GetRedisConfig() (*RedisConf, error) {
	conf := new(RedisConf)
	file, err := ioutil.ReadFile(_REDIS_CONF_PATH)
	if err != nil {
		logger.Errorf("[GetRedisConfig] open file failed. err:%v", err)
		return nil, err
	}
	if err = yaml.Unmarshal(file, conf); err != nil {
		logger.Errorf("[GetRedisConfig] unmarshal file failed. err:%v", err)
		return nil, err
	}
	return conf, nil
}
