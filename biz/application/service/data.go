package service

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/google/wire"
	"github.com/xh-polaris/platform-core-api/biz/adaptor"
	api "github.com/xh-polaris/platform-core-api/biz/application/dto/platform/core_api"
	"github.com/xh-polaris/platform-core-api/biz/infra/config"
	"github.com/xh-polaris/platform-core-api/biz/infra/rpc/platform_data"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/data"
	"time"
)

type IDataService interface {
	ReportEvent(ctx context.Context, req *api.ReportEventRequest) (*api.ReportEventResponse, error)
}

type DataService struct {
	Config       *config.Config
	PlatformData platform_data.PlatformData
}

var DataServiceSet = wire.NewSet(
	wire.Struct(new(DataService), "*"),
	wire.Bind(new(IDataService), new(*DataService)),
)

func (s *DataService) ReportEvent(ctx context.Context, req *api.ReportEventRequest) (*api.ReportEventResponse, error) {

	documents := make([]*data.Document, 0)

	resp := &api.ReportEventResponse{
		Done: false,
	}

	for _, item := range req.Data {
		eventName := item.EventName

		var tags map[string]interface{}

		err := sonic.UnmarshalString(item.Tags, &tags)

		if err != nil {
			return resp, err
		}
		// 添加时间戳
		tags["@timestamp"] = time.Now().Format(time.RFC3339)

		// 添加user meta到tags中
		userMeta := adaptor.ExtractUserMeta(ctx)
		if userMeta.UserId != "" {
			tags["@user_meta"] = userMeta
		}

		tagsString, err := sonic.MarshalString(tags)

		if err != nil {
			return resp, err
		}

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
		return resp, err
	}

	resp.Done = true

	return resp, nil
}
