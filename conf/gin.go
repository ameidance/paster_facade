package conf

import (
    "github.com/bytedance/gopkg/util/logger"
    "gopkg.in/yaml.v3"
    "io/ioutil"
)

type GinConf struct {
    Address string `yaml:"address"`
    Port    int    `yaml:"port"`
}

const (
    _GIN_CONF_PATH = "conf/gin.yml"
)

func GetGinConfig() (*GinConf, error) {
    conf := new(GinConf)
    file, err := ioutil.ReadFile(_GIN_CONF_PATH)
    if err != nil {
        logger.Errorf("[GetGinConfig] open file failed. err:%v", err)
        return nil, err
    }
    if err = yaml.Unmarshal(file, conf); err != nil {
        logger.Errorf("[GetGinConfig] unmarshal file failed. err:%v", err)
        return nil, err
    }
    return conf, nil
}
