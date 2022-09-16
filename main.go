package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"

	"github.com/nanoyou/MaidNanaGo/controller"
	_ "github.com/nanoyou/MaidNanaGo/docs"
	"github.com/nanoyou/MaidNanaGo/middleware"
	"github.com/nanoyou/MaidNanaGo/mirai"
	"github.com/nanoyou/MaidNanaGo/model"
)

//go:embed MaidNanaFrontEnd/dist/*
var frontEndDist embed.FS

// @title       Main Nana API 文档
// @version    	1.0.0-alpha
// @description Maid Nana 的 Web API

// @host	localhost:5277
// @base 	/api
func main() {
	mirai.InitBot()
	model.Init()
	app := iris.New()

	frontEnd, _ := fs.Sub(frontEndDist, "MaidNanaFrontEnd/dist")
	app.HandleDir("/", http.FS(frontEnd), iris.DirOptions{
		IndexName: "index.html",
		SPA:       true,
	})
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
