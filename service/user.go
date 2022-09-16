package service

import (
	"github.com/nanoyou/MaidNanaGo/model"
	pwd "github.com/nanoyou/MaidNanaGo/util/password"
)

type UserService struct{}

var userService = &UserService{}

func GetUserService() *UserService { return userService }

// Register 注册账号
func (u *UserService) Register(name string, password string) (user *model.User, err error) {

	user = new(model.User)
	user.Name = name
	user.HashedPassword = pwd.NewSHA512Password(password).String()

	err = user.Create()
	return
}
