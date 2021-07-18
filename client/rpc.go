package client

import (
	"github.com/ameidance/paster_facade/model/dto/kitex_gen/ameidance/paster/core/pastercoreservice"
	"github.com/cloudwego/kitex/client"
)

var (
	CoreClient pastercoreservice.Client
)

func init() {
	CoreClient = pastercoreservice.MustNewClient("ameidance.paster.core", client.WithHostPorts("0.0.0.0:8888"))
}
