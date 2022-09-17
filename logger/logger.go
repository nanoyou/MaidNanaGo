package logger

import (
	"github.com/nanoyou/MaidNanaGo/config"
	"github.com/sirupsen/logrus"
)

func Init() {
	level, err := logrus.ParseLevel(config.Config.GetString("logger.level"))
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(level)
}
