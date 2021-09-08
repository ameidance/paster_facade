package client

import (
	"github.com/ameidance/paster_facade/client/consul"
	"github.com/ameidance/paster_facade/frame"
	"github.com/ameidance/paster_facade/model/dto/kitex_gen/paster/core/core"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
)

var (
	resolver   *consul.Resolver
	options    []client.Option
	CoreClient core.Client
)

func init() {
	resolver = consul.NewResolver()
	options = append(options, client.WithClientBasicInfo(frame.EBI), client.WithResolver(resolver), client.WithTransportProtocol(transport.GRPC))
}

func InitRpc() {
	CoreClient = core.MustNewClient("ameidance.paster.core", options...)
}
