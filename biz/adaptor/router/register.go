package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"platform-core-api/biz/adaptor/router/core_api"
)

func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	core_api.Register(r)
}
