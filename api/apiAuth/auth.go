package apiauth

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func Auth() routing.Handler {
	return func(ctx *routing.Context) error {
		ctx.SetStatusCode(fasthttp.StatusAccepted)
		return ctx.Next()
	}

}
