package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	logrus.WithError(err).Fatal("无法读取文件信息")
	panic(err)
}
