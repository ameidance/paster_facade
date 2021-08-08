package client

import (
	"github.com/ameidance/paster_facade/client/consul"
	"github.com/ameidance/paster_facade/model/dto/kitex_gen/ameidance/paster/core/pastercoreservice"
	"github.com/cloudwego/kitex/client"
)

var (
	Resolver   *consul.Resolver
	CoreClient pastercoreservice.Client
)

func InitRpc() {
	Resolver = consul.NewResolver()
	CoreClient = pastercoreservice.MustNewClient("ameidance.paster.core", client.WithResolver(Resolver))
}
