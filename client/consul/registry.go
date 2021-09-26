package consul

import (
	"fmt"

	"github.com/ameidance/paster_facade/frame"
	"github.com/ameidance/paster_facade/util"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/hashicorp/consul/api"
)

type Registry struct {
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (m *Registry) Register(info *registry.Info) (err error) {
	if client == nil {
		return nil
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = frame.GetInstanceId()
	registration.Name = frame.GetServiceName()
	registration.Address = util.GetInternalIP()
	_, err = fmt.Sscanf(info.Addr.String(), "[::]:%v", &registration.Port)
	if err != nil {
		klog.Errorf("[Registry -> Register] get registry info addr port failed. err:%v", err)
		return
	}

	registration.Check = new(api.AgentServiceCheck)
	registration.Check.GRPC = fmt.Sprintf("%s:%d/%s", registration.Address, registration.Port, registration.Name)
	registration.Check.Timeout = "5s"
	registration.Check.Interval = "5s"
	registration.Check.DeregisterCriticalServiceAfter = "10s"

	klog.Infof("[Registry -> Register] registering... instance id:%v", frame.GetInstanceId())
	return client.Agent().ServiceRegister(registration)
}

func (m *Registry) Deregister(info *registry.Info) error {
	if client == nil {
		return nil
	}

	klog.Infof("[Registry -> Register] deregistering... instance id:%v", frame.GetInstanceId())
	return client.Agent().ServiceDeregister(frame.GetInstanceId())
}
