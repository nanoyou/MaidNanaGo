package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/sirupsen/logrus"

	"github.com/nanoyou/MaidNanaGo/config"
	"github.com/nanoyou/MaidNanaGo/controller"
	"github.com/nanoyou/MaidNanaGo/controller/request"
	_ "github.com/nanoyou/MaidNanaGo/docs"
	"github.com/nanoyou/MaidNanaGo/logger"
	"github.com/nanoyou/MaidNanaGo/middleware"
	"github.com/nanoyou/MaidNanaGo/mirai"
	"github.com/nanoyou/MaidNanaGo/model"
	"github.com/nanoyou/MaidNanaGo/service"
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
	// 判断是否为第一次启动
	checkFirstStart()
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

		needLogin := middleware.Auth()
		superAdmin := middleware.Role(model.SUPER_ADMIN)

		api.Get("/user", superAdmin, userController.UserList)
		api.Post("/user", userController.Register)
		user := api.Party("/user/{username}")
		{
			user.Post("/login", userController.Login)
			user.Post("/logout", needLogin, userController.Logout)
		}
	}

	app.Listen(":5277")
}

// checkFirstStart 第一次运行
func checkFirstStart() {
	// 判断是否为第一次运行
	if !config.Config.GetBool("first_start") {
		return
	}

	// 创建默认账户
	logrus.Info("首次启动, 创建默认管理员用户")
	for {
		fmt.Print("请输入用户名: ")
		var username, password, confirmPassword string
		fmt.Scan(&username)
		fmt.Print("请输入密码: ")
		fmt.Scan(&password)
		fmt.Print("请输入确认密码: ")
		fmt.Scan(&confirmPassword)
		if password != confirmPassword {
			fmt.Println("两次密码不一致")
			continue
		}

		// 校验合法性
		req := request.RegisterRequest{Username: username, Password: password, VerificationCode: 114}
		err := validator.Get().Struct(req)
		if err != nil {
			fmt.Printf("参数不合法: %v\n", err)
			continue
		}
		// 注册用户
		_, err = service.GetUserService().Register(username, password, -1)
		if err != nil {
			fmt.Printf("注册失败: %v\n", err)
			continue
		}
		// 添加管理员权限
		err = service.GetUserService().AddRole(username, model.SUPER_ADMIN)
		if err != nil {
			logrus.WithError(err).Fatal("添加权限失败")
			panic(err)
		}
		config.Config.Set("first_start", false)
		err = config.Config.WriteConfig()
		if err != nil {
			logrus.WithError(err).Fatal("写出配置文件失败")
			panic(err)
		}
		fmt.Println("注册成功")
		return
	}
}
