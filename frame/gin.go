package frame

import (
	"fmt"
	"io/ioutil"

	"github.com/ameidance/paster_facade/constant"
	"github.com/cloudwego/kitex/pkg/klog"
	"gopkg.in/yaml.v3"
)

type _GinConf struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

var (
	GinConf *_GinConf
)

func init() {
	var err error
	if GinConf, err = getGinConfig(); err != nil {
		panic(err)
	}
}

func GetGinAddress() string {
	if GinConf == nil {
		return ""
	}
	return fmt.Sprintf("%s:%d", GinConf.Address, GinConf.Port)
}

func getGinConfig() (*_GinConf, error) {
	conf := new(_GinConf)
	file, err := ioutil.ReadFile(constant.GIN_CONF_PATH)
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
