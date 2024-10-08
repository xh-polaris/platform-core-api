package platform_data

import (
	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/kitex/client"
	"github.com/xh-polaris/platform-core-api/biz/infra/config"
	data "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/data/dataservice"
)

type IPlatformData interface {
	data.Client
}

type PlatformData struct {
	data.Client
}

var PlatformDataSet = wire.NewSet(
	NewPlatformData,
	wire.Struct(new(PlatformData), "*"),
	wire.Bind(new(IPlatformData), new(*PlatformData)),
)

func NewPlatformData(config *config.Config) data.Client {
	return client.NewClient(config.Name, "platform.data", data.NewClient)
}
