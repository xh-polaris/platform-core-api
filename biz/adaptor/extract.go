package adaptor

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

const hertzContext = "hertz_context"

func InjectContext(ctx context.Context, c *app.RequestContext) context.Context {
	return context.WithValue(ctx, hertzContext, c)
}
