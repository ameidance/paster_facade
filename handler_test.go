package main

import (
	"context"
	"testing"

	"github.com/ameidance/paster_facade/client"
	"github.com/ameidance/paster_facade/client/consul"
	"github.com/ameidance/paster_facade/model/vo/kitex_gen/facade"
	"github.com/ameidance/paster_facade/util"
	"github.com/cloudwego/kitex/pkg/klog"
)

func TestSavePostHandler(*testing.T) {
	client.InitRedis()
	consul.InitConsul()
	client.InitRpc()

	ctx := context.WithValue(context.Background(), "ip", util.GetInternalIP())
	resp, _ := new(PasterFacadeImpl).SavePost(ctx, &facade.SavePostRequest{
		Content:      "ok",
		Lang:         1,
		Nickname:     "2U",
		IsDisposable: false,
		Passwd:       "",
	})
	klog.Infof("[TestSavePostHandler] resp:%v", util.GetJsonString(resp))
}

func TestGetPostHandler(*testing.T) {
	client.InitRedis()
	consul.InitConsul()
	client.InitRpc()

	resp, _ := new(PasterFacadeImpl).GetPost(context.Background(), &facade.GetPostRequest{
		Id:     48,
		Passwd: "",
	})
	klog.Infof("[TestGetPostHandler] resp:%v", util.GetJsonString(resp))
}

func TestGetCommentsHandler(*testing.T) {
	client.InitRedis()
	consul.InitConsul()
	client.InitRpc()

	resp, _ := new(PasterFacadeImpl).GetComments(context.Background(), &facade.GetCommentsRequest{
		PostId: 1,
		Passwd: "",
	})
	klog.Infof("[TestGetCommentsHandler] resp:%v", util.GetJsonString(resp))
}
