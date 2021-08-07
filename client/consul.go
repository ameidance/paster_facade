package client

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/ameidance/paster_facade/constant"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v3"
)

var (
	ConsulClient *api.Client
)

type _ConsulConf struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
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

func InitConsul() {
	address, err := getConsulConfig()
	if address == nil || err != nil {
		panic(err)
	}
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%v:%v", address.Hostname, address.Port)
	ConsulClient, err = api.NewClient(config)
	if ConsulClient == nil || err != nil {
		panic(err)
	}
}

type ConsulResolver struct {
}

func NewConsulResolver() *ConsulResolver {
	return &ConsulResolver{}
}

func (*ConsulResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	if ConsulClient == nil {
		return
	}
	klog.Infof("[ConsulResolver -> Target] service name:%v", target.ServiceName())
	return target.ServiceName()
}

func (*ConsulResolver) Resolve(ctx context.Context, desc string) (res discovery.Result, err error) {
	if ConsulClient == nil {
		return
	}
	var services []*api.ServiceEntry
	services, _, err = ConsulClient.Health().Service(desc, "", true, nil)
	if err != nil {
		klog.Errorf("[ConsulResolver -> Resolve] find healthy services failed. target service name:%v", desc)
		return
	}
	for _, service := range services {
		addr := fmt.Sprintf("%v:%v", service.Service.Address, service.Service.Port)
		res.Instances = append(res.Instances, discovery.NewInstance("tcp", addr, 10, nil))
	}
	res.Cacheable = true
	res.CacheKey = desc
	return
}

func (*ConsulResolver) Diff(cacheKey string, prev, next discovery.Result) (discovery.Change, bool) {
	return discovery.DefaultDiff(cacheKey, prev, next)
}

func (*ConsulResolver) Name() string {
	return "ConsulResolver"
}
