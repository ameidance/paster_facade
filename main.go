package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ameidance/paster_facade/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

var router *gin.Engine

func main() {
	client.InitRedis()

	ginConf, err := getGinConfig()
	if err != nil {
		panic(err)
	}
	err = router.Run(fmt.Sprintf("%s:%d", ginConf.Address, ginConf.Port))
	if err != nil {
		panic(err)
	}
}

type _GinConf struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

const (
	_GIN_CONF_PATH = "conf/gin.yml"
)

func getGinConfig() (*_GinConf, error) {
	conf := new(_GinConf)
	file, err := ioutil.ReadFile(_GIN_CONF_PATH)
	if err != nil {
		klog.Errorf("[getGinConfig] open file failed. err:%v", err)
		return nil, err
	}
	if err = yaml.Unmarshal(file, conf); err != nil {
		klog.Errorf("[getGinConfig] unmarshal file failed. err:%v", err)
		return nil, err
	}
	return conf, nil
}
