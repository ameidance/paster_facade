package main

import (
	"github.com/ameidance/paster_facade/client"
	"github.com/ameidance/paster_facade/client/consul"
	"github.com/ameidance/paster_facade/frame"
	"github.com/ameidance/paster_facade/model/vo/kitex_gen/facade/pasterfacade"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func main() {
	client.InitRedis()
	consul.InitConsul()
	client.InitRpc()

	srv := pasterfacade.NewServer(new(PasterFacadeImpl), server.WithServiceAddr(frame.Address),
		server.WithServerBasicInfo(frame.EBI), server.WithRegistry(consul.NewRegistry()))
	if err := srv.Run(); err != nil {
		klog.Errorf("[main] server stopped with error. err:%v", err)
		panic(err)
	} else {
		klog.Infof("[main] server stopped.")
	}
}
