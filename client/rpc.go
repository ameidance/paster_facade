package client

import (
	"github.com/ameidance/paster_facade/model/dto/kitex_gen/ameidance/paster/core/pastercoreservice"
	"github.com/cloudwego/kitex/client"
)

var (
	Resolver   *ConsulResolver
	CoreClient pastercoreservice.Client
)

func InitRpc() {
	Resolver = NewConsulResolver()
	CoreClient = pastercoreservice.MustNewClient("ameidance.paster.core", client.WithResolver(Resolver))
}
