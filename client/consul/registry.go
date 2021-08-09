package consul

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ameidance/paster_facade/frame"
	"github.com/ameidance/paster_facade/util"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/hashicorp/consul/api"
)

var (
	CheckCounter int64
)

type Registry struct {
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (m *Registry) Register(info *registry.Info) (err error) {
	if Client == nil {
		return nil
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = frame.GetInstanceId()
	registration.Name = frame.GetServiceName()
	registration.Address = util.GetInternalIP()
	_, err = fmt.Sscanf(info.Addr.String(), ":%v", &registration.Port)
	if err != nil {
		klog.Errorf("[Registry -> Register] get registry info addr port failed. err:%v", err)
		return
	}

	registration.Check = new(api.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%d/health", registration.Address, registration.Port)
	registration.Check.Timeout = "5s"
	registration.Check.Interval = "5s"
	registration.Check.DeregisterCriticalServiceAfter = "10s"

	klog.Infof("[Registry -> Register] registering... instance id:%v", frame.GetInstanceId())
	return Client.Agent().ServiceRegister(registration)
}

func (m *Registry) Deregister(info *registry.Info) error {
	if Client == nil {
		return nil
	}

	klog.Infof("[Registry -> Register] deregistering... instance id:%v", frame.GetInstanceId())
	return Client.Agent().ServiceDeregister(frame.GetInstanceId())
}

func (m *Registry) Initialize() (err error) {
	info := new(registry.Info)
	info.Addr = utils.NewNetAddr("tcp", fmt.Sprintf(":%d", frame.GinConf.Port))
	if err = m.Register(info); err != nil {
		return
	}
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, os.Kill)
		select {
		case <-ch:
			_ = m.Deregister(info)
			os.Exit(0)
		}
	}()
	return
}
