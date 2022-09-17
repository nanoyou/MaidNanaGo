package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	logrus.Info("加载配置文件")
	Config = viper.New()
	Config.SetConfigFile("config.json")
	Config.SetConfigType("json")

	// Config.SetDefault("bot.qq", 114514191810)
	// Config.SetDefault("bot.password", "QQ密码")
	Config.SetDefault("bot.loginmethod", "common")

	Config.SetDefault("logger.level", logrus.InfoLevel)

	err := Config.ReadInConfig()
	if err != nil {
		writeDefaultConfig()
	}
	logrus.Info("配置文件加载完成")
}

func writeDefaultConfig() {
	logrus.Info("正在写出默认配置文件")
	err := Config.WriteConfig()
	if err != nil {
		logrus.WithError(err).Fatal("无法写出默认配置文件")
		os.Exit(1)
	}
}
