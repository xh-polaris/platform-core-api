package core_api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"platform-core-api/biz/adaptor"
	"platform-core-api/biz/common/request"
	"platform-core-api/provider"
)

// ReportEvent
// @router /platform/report_event [POST]
func ReportEvent(ctx context.Context, c *app.RequestContext) {
	var err error
	var req request.InsertRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	p := provider.Get()
	resp, err := p.DataService.ReportEvent(ctx, &req)
	adaptor.PostProcess(ctx, c, &req, resp, err)
}
