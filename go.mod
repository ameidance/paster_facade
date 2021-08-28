module github.com/ameidance/paster_facade

go 1.15

require (
	github.com/cloudwego/kitex v0.0.4
	github.com/go-redis/redis/v8 v8.11.0
	github.com/hashicorp/consul/api v1.9.1
	github.com/json-iterator/go v1.1.11
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
