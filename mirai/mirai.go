package mirai

import (
	"os"

	_ "github.com/nanoyou/MaidNanaGo/mirai/module/file"
	_ "github.com/nanoyou/MaidNanaGo/mirai/module/verification"

	"github.com/Logiase/MiraiGo-Template/bot"
	lc "github.com/Logiase/MiraiGo-Template/config"
	"github.com/nanoyou/MaidNanaGo/config"
	"github.com/nanoyou/MaidNanaGo/util/file"
	"github.com/sirupsen/logrus"
)

func InitBot() {
	lc.GlobalConfig = &lc.Config{Viper: config.Config}

	bot.InitBot(config.Config.GetInt64("bot.qq"), config.Config.GetString("bot.password"))
	if !file.FileExists("./device.json") {
		bot.GenRandomDevice()
	}
	bot.StartService()
	deviceJSON, err := os.ReadFile("./device.json")
	if err != nil {
		logrus.WithError(err).Fatal("无法读取设备文件")
		panic(err)
	}
	bot.UseDevice(deviceJSON)
	err = bot.Login()
	if err != nil {
		logrus.WithError(err).Fatal("无法登陆")
	}
	bot.SaveToken()

}
