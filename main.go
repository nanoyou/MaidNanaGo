package main

import (
	"embed"
	"io/fs"
	"net/http"
	"time"

	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"

	"github.com/nanoyou/MaidNanaGo/controller"
	_ "github.com/nanoyou/MaidNanaGo/docs"
	"github.com/nanoyou/MaidNanaGo/logger"
	"github.com/nanoyou/MaidNanaGo/middleware"
	"github.com/nanoyou/MaidNanaGo/mirai"
	"github.com/nanoyou/MaidNanaGo/model"
	"github.com/nanoyou/MaidNanaGo/validator"
)

//go:embed MaidNanaFrontEnd/dist/*
var frontEndDist embed.FS

// @title       Main Nana API 文档
// @version    	1.0.0-alpha
// @description Maid Nana 的 Web API

// @host	localhost:5277
// @base 	/api
func main() {
	// 初始化日志
	logger.Init()
	// 初始化 QQ 机器人
	mirai.InitBot()
	// 初始化数据库
	model.Init()
	// 初始化 web 框架
	app := iris.New()
	// 注入校验器
	app.Validator = validator.Get()

	session := sessions.New(sessions.Config{
		Cookie: "MaidNana",
		// Expires: time.Hour * 24 * 7,
		Expires: time.Minute, //测试
	})
	app.Use(session.Handler())

	// 处理静态文件
	frontEnd, _ := fs.Sub(frontEndDist, "MaidNanaFrontEnd/dist")
	app.HandleDir("/", http.FS(frontEnd), iris.DirOptions{
		IndexName: "index.html",
		SPA:       true,
	})

	api := app.Party("/api")
	{
		// 跨域中间件
		api.Use(middleware.Cors())

		// API 文档
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

		// Controller
		debugController := new(controller.DebugController)
		api.Get("/about", debugController.About)
		userController := new(controller.UserController)

		api.Post("/user", userController.Register)
		user := api.Party("/user/{username}")
		{
			user.Post("/login", userController.Login)
		}
	}

	app.Listen(":5277")
}
