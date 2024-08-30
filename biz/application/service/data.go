package service

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/data"
	api "platform-core-api/biz/application/dto/platform/core_api"
	"platform-core-api/biz/infra/config"
	"platform-core-api/biz/infra/rpc/platform_data"
	"time"
)

type IDataService interface {
	ReportEvent(ctx context.Context, req *api.ReportEventRequest) (bool, error)
}

type DataService struct {
	Config       *config.Config
	PlatformData platform_data.PlatformData
}

var DataServiceSet = wire.NewSet(
	wire.Struct(new(DataService), "*"),
	wire.Bind(new(IDataService), new(*DataService)),
)

func (s *DataService) ReportEvent(ctx context.Context, req *api.ReportEventRequest) (bool, error) {

	documents := make([]*data.Document, 0)

	for _, item := range req.Data {
		eventName := item.EventName

		var tags map[string]interface{}

		err := sonic.UnmarshalString(item.Tags, &tags)

		if err != nil {
			return false, err
		}
		// 添加时间戳
		tags["@timestamp"] = time.Now().Format(time.RFC3339)

		// TODO 添加user meta到tags中

		tagsString, err := sonic.MarshalString(tags)

		if err != nil {
			return false, err
		}

		fmt.Println(tagsString)

		document := data.Document{
			EventName: eventName,
			Tags:      tagsString,
		}

		documents = append(documents, &document)
	}

	insertReq := &data.InsertReq{
		Documents: documents,
	}

	_, err := s.PlatformData.Insert(ctx, insertReq)

	if err != nil {
		return false, err
	}
	return true, nil
}
