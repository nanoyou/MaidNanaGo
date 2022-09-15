package main

import (
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"

	"github.com/nanoyou/MaidNanaGo/controller"
	_ "github.com/nanoyou/MaidNanaGo/docs"
	"github.com/nanoyou/MaidNanaGo/middleware"
	"github.com/nanoyou/MaidNanaGo/mirai"
)

// @title       Main Nana API 文档
// @version    	1.0.0-alpha
// @description Maid Nana 的 Web API

// @host	localhost:5277
// @base 	/api
func main() {
	mirai.InitBot()
	app := iris.New()

	api := app.Party("/api")
	{
		api.Use(middleware.Cors())
		swaggerConfig := swagger.Config{
			URL:          "http://localhost:5277/api/swagger/doc.json",
			DeepLinking:  true,
			DocExpansion: "list",
			DomID:        "#swagger-ui",
			Prefix:       "/api/swagger",
		}
		swaggerUI := swagger.Handler(swaggerFiles.Handler, swaggerConfig)
		api.Get("/swagger", swaggerUI)
		api.Get("/swagger/{any:path}", swaggerUI)

		debugController := new(controller.DebugController)
		api.Get("/about", debugController.About)
	}

	app.Listen(":5277")
}
