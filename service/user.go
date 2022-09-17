package service

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/coocood/freecache"
	"github.com/nanoyou/MaidNanaGo/model"
	"github.com/nanoyou/MaidNanaGo/util/bytes"
	pwd "github.com/nanoyou/MaidNanaGo/util/password"
	"github.com/sirupsen/logrus"
)

type UserService struct{}

var (
	verificationCode = freecache.NewCache(1024 * 1024 * 4)
)

var userService = &UserService{}

func GetUserService() *UserService { return userService }

// Register 注册账号
func (u *UserService) Register(name string, password string, qq int64) (user *model.User, err error) {

	user = new(model.User)
	user.Name = name
	user.QQ = qq
	user.HashedPassword = pwd.NewSHA512Password(password).String()

	err = user.Create()
	return
}

// CreateVerificationCode 创建 QQ 验证码, 过期时间为10分钟
func (u *UserService) CreateVerificationCode(qq int64) (code int, err error) {
	user, err := model.GetUserByQQ(qq)
	if err == nil {
		return 0, fmt.Errorf("该 QQ 号已被 %v 绑定", user.Name)
	}
	for {
		code = rand.Intn(900000) + 100000
		logrus.WithField("code", code).Debug("生成验证码")
		if _, err := verificationCode.GetInt(int64(code)); err != nil {
			if err == freecache.ErrNotFound {
				break
			}
			logrus.WithError(err).Fatal("创建验证码时读取缓存出错")
			return 0, err
		}
	}
	err = verificationCode.SetInt(int64(code), bytes.Int64ToBytes(qq), 60*10)
	if err != nil {
		logrus.WithError(err).Fatal("创建验证码时写入缓存出错")
		return 0, err
	}

	return code, nil
}

// GetQQByVerificationCode 通过验证码获取 QQ 号
func (u *UserService) GetQQByVerificationCode(code int) (qq int64, err error) {
	var qqBytes []byte
	if qqBytes, err = verificationCode.GetInt(int64(code)); err != nil {
		if errors.Is(err, freecache.ErrNotFound) {
			return 0, err
		}
		logrus.WithError(err).Fatal("获取QQ号时读取缓存出错")
		return 0, err
	}
	return bytes.BytesToInt64(qqBytes), nil
}
