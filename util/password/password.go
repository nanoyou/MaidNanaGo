package password

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type Password interface {
	// Validate 判断密码是否正确
	Validate(password string) bool
	// String 将密码内容转成字符串
	String() string
}

func FromString(cipherText string) Password {
	arr := strings.Split(cipherText, ":")
	if len(arr) != 3 {
		logrus.Panic("密码格式错误")
		return nil
	}

	method := arr[0]
	salt := arr[1]
	text := arr[2]

	switch method {
	case "PLAIN":
		return &PlainPassword{text}
	case "SHA-512":
		return &SHA512Password{text, salt}
	}
	logrus.Panic("密码格式错误")
	return nil
}
