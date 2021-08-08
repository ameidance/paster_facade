package consul

import (
	"fmt"
	"io/ioutil"

	"github.com/ameidance/paster_facade/constant"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v3"
)

var (
	Config *_ConsulConf
	Client *api.Client
)

type _ConsulConf struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}

func InitConsul(servicePort int) {
	var err error
	Config, err = getConsulConfig()
	if Config == nil || err != nil {
		panic(err)
	}
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%v:%v", Config.Hostname, Config.Port)
	Client, err = api.NewClient(config)
	if Client == nil || err != nil {
		panic(err)
	}
	if NewRegistry().Initialize(servicePort) != nil {
		panic(err)
	}
}

func getConsulConfig() (*_ConsulConf, error) {
	conf := new(_ConsulConf)
	file, err := ioutil.ReadFile(constant.CONSUL_CONF_PATH)
	if err != nil {
		klog.Errorf("[getConsulConfig] open file failed. err:%v", err)
		return nil, err
	}
	if err = yaml.Unmarshal(file, conf); err != nil {
		klog.Errorf("[getConsulConfig] unmarshal file failed. err:%v", err)
		return nil, err
	}
	return conf, nil
}
