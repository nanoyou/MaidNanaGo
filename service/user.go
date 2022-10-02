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
	// 验证码限流缓存
	verificationCodeLimit = freecache.NewCache(1024 * 1024 * 1)
	// 验证码缓存
	verificationCode = freecache.NewCache(1024 * 1024 * 4)
)

var userService = &UserService{}

func GetUserService() *UserService { return userService }

// Register 注册账号
func (u *UserService) Register(name string, password string, qq int64) (user *model.User, err error) {

	user = new(model.User)
	user.Name = name
	user.QQ = qq
	// 散列密码
	user.HashedPassword = pwd.NewSHA512Password(password).String()

	// 写入数据库
	err = user.Create()
	return
}

// Login 登录
func (u *UserService) Login(name string, password string) (user *model.User, err error) {
	// 查找用户是否存在
	user, err = model.GetUserByName(name)
	if err != nil {
		return nil, errors.New("用户名不存在")
	}

	// 验证密码是否正确
	if !pwd.FromString(user.HashedPassword).Validate(password) {
		return nil, errors.New("密码错误")
	}
	logrus.WithField("user", name).Debug("登录")

	return user, nil
}

// CreateVerificationCode 创建 QQ 验证码, 过期时间为10分钟
func (u *UserService) CreateVerificationCode(qq int64) (code int, err error) {
	// 检查该 QQ 是否已经被绑定
	user, err := model.GetUserByQQ(qq)
	if err == nil {
		return 0, fmt.Errorf("该 QQ 号已被 %v 绑定", user.Name)
	}

	// 检查是否已经获取过验证码
	if _, err = verificationCodeLimit.GetInt(qq); err == nil {
		return 0, fmt.Errorf("请求过快, 请 1 分钟后重试")
	} else if err != freecache.ErrNotFound {
		logrus.WithError(err).Fatal("创建验证码时读取缓存出错")
		return 0, err
	}

	// 生成验证码
	for {
		code = rand.Intn(900000) + 100000
		logrus.WithField("code", code).Debug("生成验证码")

		// 判断验证码是否与现有验证码冲突
		if _, err := verificationCode.GetInt(int64(code)); err != nil {
			if err != freecache.ErrNotFound {
				logrus.WithError(err).Fatal("创建验证码时读取缓存出错")
				return 0, err
			}
			break
		}
	}
	// 将(验证码 -> QQ)写入缓存
	err = verificationCode.SetInt(int64(code), bytes.Int64ToBytes(qq), 60*10)
	if err != nil {
		logrus.WithError(err).Fatal("创建验证码时写入缓存出错")
		return 0, err
	}

	// 标记 QQ 号
	err = verificationCodeLimit.SetInt(qq, []byte{}, 60)
	if err != nil {
		logrus.WithError(err).Fatal("创建验证码时写入缓存出错")
		return 0, err
	}

	return code, nil
}

// GetQQByVerificationCode 通过验证码获取 QQ 号
func (u *UserService) GetQQByVerificationCode(code int) (qq int64, err error) {
	var qqBytes []byte
	// 在缓存中查找验证码对应的QQ号
	if qqBytes, err = verificationCode.GetInt(int64(code)); err != nil {
		if errors.Is(err, freecache.ErrNotFound) {
			return 0, err
		}
		logrus.WithError(err).Fatal("获取QQ号时读取缓存出错")
		return 0, err
	}
	return bytes.BytesToInt64(qqBytes), nil
}

// SetRole 设置用户权限, 如果权限是管理员则删除所有原有权限
func (u *UserService) SetRole(username string, role model.RoleType) error {
	// 查找用户
	user, err := model.GetUserByName(username)
	if err != nil {
		return err
	}

	roles := []model.RoleType{role}
	for _, oldRole := range user.Roles {
		if oldRole.Role != role {
			roles = append(roles, oldRole.Role)
		}
	}
	logrus.WithField("roles", roles).Debug("将要添加权限")

	// 如果包含管理员角色 设置仅设置为管理员
	for _, oldRole := range roles {
		if oldRole == model.SUPER_ADMIN {
			roles = []model.RoleType{model.SUPER_ADMIN}
			break
		}
	}
	logrus.WithField("user", username).WithField("roles", roles).Debug("设置权限")

	// 保存角色
	return user.SetRole(roles)

}

// DeleteRole 删除用户权限
func (u *UserService) DeleteRole(username string, role model.RoleType) error {
	// 查找用户
	user, err := model.GetUserByName(username)
	if err != nil {
		return err
	}

	roles := []model.RoleType{}
	for _, oldRole := range user.Roles {
		if oldRole.Role != role {
			roles = append(roles, oldRole.Role)
		}
	}

	logrus.WithField("user", username).WithField("roles", roles).Debug("删除权限")

	// 保存角色
	return user.SetRole(roles)

}

// GetAllUsers 获取全部用户
func (u *UserService) GetAllUsers() ([]model.User, error) {
	return model.GetAllUsers()
}

// GetUser 获取用户
func (u *UserService) GetUser(username string) (*model.User, error) {
	return model.GetUserByName(username)
}

// ModifyUser 修改用户信息
func (u *UserService) ModifyUser(user *model.User) error {
	return user.Update()
}

// ChangePassword 修改用户密码
func (u *UserService) ChangePassword(username string, password string) error {
	user, err := u.GetUser(username)
	if err != nil {
		return err
	}

	// 散列密码
	user.HashedPassword = pwd.NewSHA512Password(password).String()
	return user.Update()
}

// DeleteUser 删除用户
func (u *UserService) DeleteUser(username string) (*model.User, error) {
	// TODO: implement
	// 先获取再删除
	return nil, nil
}
