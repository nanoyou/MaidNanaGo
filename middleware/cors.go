package middleware

import (
	"github.com/kataras/iris/v12/context"
)

func Cors() context.Handler {
	return func(ctx *context.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Next()
	}
}
