package main

import (
	"github.com/ameidance/paster_facade/client"
	"github.com/ameidance/paster_facade/client/consul"
	"github.com/ameidance/paster_facade/frame"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	client.InitRedis()
	consul.InitConsul()
	client.InitRpc()

	if err := router.Run(frame.GetGinAddress()); err != nil {
		klog.Errorf("[main] server stopped with error. err:%v", err)
		panic(err)
	}
}
