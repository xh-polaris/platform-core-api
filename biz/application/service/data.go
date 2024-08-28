package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/data"
	"platform-core-api/biz/common/request"
	"platform-core-api/biz/infra/config"
	"platform-core-api/biz/infra/rpc/platform_data"
	"strings"
)

type IDataService interface {
	ReportEvent(ctx context.Context, req *request.InsertRequest) (bool, error)
}

type DataService struct {
	Config       *config.Config
	PlatformData platform_data.PlatformData
}

var DataServiceSet = wire.NewSet(
	wire.Struct(new(DataService), "*"),
	wire.Bind(new(IDataService), new(*DataService)),
)

func (s *DataService) ReportEvent(ctx context.Context, req *request.InsertRequest) (bool, error) {

	documents := make([]*data.Document, 0)

	for _, item := range req.Body {
		eventName := item.EventName
		tags := item.Tags

		tagsString, err := json.Marshal(tags)
		tagsString = []byte(strings.ReplaceAll(string(tagsString), "\"", "'"))

		fmt.Println(string(tagsString))

		if err != nil {
			return false, err
		}

		document := data.Document{
			EventName: eventName,
			Tags:      string(tagsString),
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
