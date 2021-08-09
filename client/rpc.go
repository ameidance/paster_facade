package client

import (
	"github.com/ameidance/paster_facade/client/consul"
	"github.com/ameidance/paster_facade/frame"
	"github.com/ameidance/paster_facade/model/dto/kitex_gen/ameidance/paster/core/pastercoreservice"
	"github.com/cloudwego/kitex/client"
)

var (
	resolver   *consul.Resolver
	options    []client.Option
	CoreClient pastercoreservice.Client
)

func init() {
	resolver = consul.NewResolver()
	options = append(options, client.WithClientBasicInfo(frame.EBI), client.WithResolver(resolver), client.WithMiddleware(frame.LogMiddleware))
}

func InitRpc() {
	CoreClient = pastercoreservice.MustNewClient("ameidance.paster.core", options...)
}
