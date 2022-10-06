package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
	"github.com/nanoyou/MaidNanaGo/model"
	"github.com/nanoyou/MaidNanaGo/util/slice"
	"github.com/sirupsen/logrus"
)

func getUser(ctx *context.Context) (user *model.User, ok bool) {
	session := sessions.Get(ctx)
	username := session.GetString("user")
	if username == "" {
		logrus.WithField("session", session.GetAll()).Debug("session")
		ctx.JSON(iris.Map{
			"ok":            false,
			"error_message": "尚未登录",
		})
		ctx.StopExecution()
		return nil, false
	}
	user, err := model.GetUserByName(username)
	if err != nil {
		ctx.JSON(iris.Map{
			"ok":            false,
			"error_message": "登录用户不存在",
			"error":         err.Error(),
		})
		ctx.StopExecution()
		return nil, false
	}
	ctx.Values().Set("user", user)
	return user, true
}

// Role 用户权限校验中间件
func Role(role model.RoleType) context.Handler {
	return func(ctx *context.Context) {
		// 获取 session 中的用户
		user, ok := getUser(ctx)
		if !ok {
			return
		}
		roles := slice.Map(user.Roles, func(rt model.Role) model.RoleType { return rt.Role })
		// 如果用户是超级管理员 直接放行
		if slice.Contains(roles, model.SUPER_ADMIN) {
			ctx.Next()
			return
		}

		if slice.Contains(roles, role) {
			ctx.Next()
			return
		}
		ctx.JSON(iris.Map{
			"ok":            false,
			"error_message": "无权访问",
			"error":         "需要权限 " + role + "",
		})
		ctx.StopExecution()

	}
}

// Auth 用户登录校验中间件
func Auth() context.Handler {
	return func(ctx *context.Context) {
		// 获取 session 中的用户
		_, ok := getUser(ctx)
		if !ok {
			return
		}

		ctx.Next()
	}
}
