package main

import (
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"

	_ "github.com/nanoyou/MaidNanaGo/docs"
	"github.com/nanoyou/MaidNanaGo/middleware"
	"github.com/nanoyou/MaidNanaGo/mirai"
)

func main() {
	mirai.InitBot()
	app := iris.New()

	api := app.Party("/api")
	{
		api.Use(middleware.Cors())
		swaggerConfig := swagger.Config{
			URL:          "http://127.0.0.1:5277/api/swagger/doc.json",
			DeepLinking:  true,
			DocExpansion: "list",
			DomID:        "#swagger-ui",
			Prefix:       "/api/swagger",
		}
		swaggerUI := swagger.Handler(swaggerFiles.Handler, swaggerConfig)
		api.Get("/swagger", swaggerUI)
		api.Get("/swagger/{any:path}", swaggerUI)
	}
	app.Listen(":5277")
}
