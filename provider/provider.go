package provider

import (
	"github.com/google/wire"
	"platform-core-api/biz/application/service"
	"platform-core-api/biz/infra/config"
	"platform-core-api/biz/infra/rpc/platform_data"
	"platform-core-api/biz/infra/rpc/platform_sts"
)

var provider *Provider

func Init() {
	var err error
	provider, err = NewProvider()
	if err != nil {
		panic(err)
	}
}

// Provider 提供controller依赖的对象
type Provider struct {
	Config      *config.Config
	DataService service.IDataService
	AuthService service.IAuthService
}

func Get() *Provider {
	return provider
}

var RPCSet = wire.NewSet(
	platform_data.PlatformDataSet,
	platform_sts.PlatformStsSet,
)

var ApplicationSet = wire.NewSet(
	service.DataServiceSet,
	service.AuthServiceSet,
)

var DomainSet = wire.NewSet()

var InfrastructureSet = wire.NewSet(
	config.NewConfig,
	RPCSet,
)

var AllProvider = wire.NewSet(
	ApplicationSet,
	DomainSet,
	InfrastructureSet,
)
