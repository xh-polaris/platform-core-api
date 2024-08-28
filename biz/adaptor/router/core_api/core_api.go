package core_api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"platform-core-api/biz/adaptor/controller/core_api"
)

func Register(r *server.Hertz) {
	root := r.Group("/", rootMw()...)
	{
		_platform := root.Group("/platform", _platformMw()...)
		_platform.POST("/report_event", append(_reporteventMw(), core_api.ReportEvent)...)
	}
}
