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
	_config *_ConsulConf
	Client  *api.Client
)

type _ConsulConf struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}

func InitConsul() {
	var err error
	_config, err = getConsulConfig()
	if _config == nil || err != nil {
		panic(err)
	}
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%v:%v", _config.Hostname, _config.Port)
	Client, err = api.NewClient(config)
	if Client == nil || err != nil {
		panic(err)
	}
	if NewRegistry().Initialize() != nil {
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
