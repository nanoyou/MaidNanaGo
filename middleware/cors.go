package middleware

import (
	"github.com/kataras/iris/v12/context"
)

func Cors() context.Handler {
	return func(ctx *context.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		if ctx.Request().Method == "OPTIONS" {
			ctx.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,HEAD,PUT,PATCH")
			ctx.Header("Access-Control-Allow-Headers", "Content-Type,Accept,Authorization")
			ctx.StatusCode(204)
			return
		}
		ctx.Next()
	}
}
