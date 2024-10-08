package platform_sts

import (
	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/kitex/client"
	sts "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts/stsservice"
	"platform-core-api/biz/infra/config"
)

type IPlatformSts interface {
	sts.Client
}

type PlatformSts struct {
	sts.Client
}

var PlatformStsSet = wire.NewSet(
	NewPlatformSts,
	wire.Struct(new(PlatformSts), "*"),
	wire.Bind(new(IPlatformSts), new(*PlatformSts)),
)

func NewPlatformSts(config *config.Config) sts.Client {
	return client.NewClient(config.Name, "platform.sts", sts.NewClient)
}