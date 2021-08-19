package consul

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/hashicorp/consul/api"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (*Resolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	if client == nil {
		return
	}

	return target.ServiceName()
}

func (*Resolver) Resolve(ctx context.Context, desc string) (res discovery.Result, err error) {
	if client == nil {
		return
	}

	var services []*api.ServiceEntry
	services, _, err = client.Health().Service(desc, "", true, nil)
	if err != nil {
		klog.Errorf("[Resolver -> Resolve] find healthy services failed. target service name:%v", desc)
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

func (*Resolver) Diff(cacheKey string, prev, next discovery.Result) (discovery.Change, bool) {
	return discovery.DefaultDiff(cacheKey, prev, next)
}

func (*Resolver) Name() string {
	return "ConsulResolver"
}
