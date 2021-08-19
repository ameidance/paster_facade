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
	config *consulConf
	client *api.Client
)

type consulConf struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}

func InitConsul() {
	var err error
	config, err = getConsulConfig()
	if config == nil || err != nil {
		panic(err)
	}
	apiConfig := api.DefaultConfig()
	apiConfig.Address = fmt.Sprintf("%v:%v", config.Hostname, config.Port)
	client, err = api.NewClient(apiConfig)
	if client == nil || err != nil {
		panic(err)
	}
}

func getConsulConfig() (*consulConf, error) {
	conf := new(consulConf)
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
