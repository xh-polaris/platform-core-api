// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"platform-core-api/biz/application/service"
	"platform-core-api/biz/infra/config"
	"platform-core-api/biz/infra/rpc/platform_data"
)

// Injectors from wire.go:

func NewProvider() (*Provider, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	client := platform_data.NewPlatformData(configConfig)
	platformData := platform_data.PlatformData{
		Client: client,
	}
	dataService := &service.DataService{
		Config:       configConfig,
		PlatformData: platformData,
	}
	providerProvider := &Provider{
		Config:      configConfig,
		DataService: dataService,
	}
	return providerProvider, nil
}
