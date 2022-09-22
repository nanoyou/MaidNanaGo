package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/nanoyou/MaidNanaGo/controller/request"
	"github.com/nanoyou/MaidNanaGo/controller/response"
	"github.com/nanoyou/MaidNanaGo/service"
)

type UserController struct{}

// @summary 		注册
// @description	 	注册账号, 需要先在 QQ 上私聊机器人发送 "验证码" 获取验证码
// @accept 			json
// @produce 		json
// @param			body body request.RegisterRequest true "注册参数"
// @tags			user
// @router 			/api/user [post]
// @success 		200	{object} response.UserResponse
// @failure 		200	{object} response.FailureResponse
func (uc *UserController) Register(ctx iris.Context) {
	// 读取 http 参数体
	var body request.RegisterRequest
	err := ctx.ReadJSON(&body)
	if err != nil {
		// 参数不合法
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "参数错误"
		ctx.JSON(r)
		return
	}

	// 获取验证码对应的 QQ 号
	qq, err := service.GetUserService().GetQQByVerificationCode(body.VerificationCode)
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "验证码无效"
		ctx.JSON(r)
		return
	}

	// 注册用户
	user, err := service.GetUserService().Register(body.Username, body.Password, qq)
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "注册失败"
		ctx.JSON(r)
		return
	}
	r := &response.UserResponse{}
	r.Ok = true
	r.User = user
	ctx.JSON(r)

}

// @summary 		登录
// @description	 	登录账号
// @accept 			json
// @produce 		json
// @param 			username path string true "用户名"
// @param			body body request.LoginRequest true "登录参数"
// @tags			user
// @router 			/api/user/{username}/login [post]
// @success 		200	{object} response.UserResponse
// @failure 		200	{object} response.FailureResponse
func (uc *UserController) Login(ctx iris.Context) {
	// 读取 http 参数体
	var body request.LoginRequest
	err := ctx.ReadJSON(&body)
	if err != nil {
		// 参数不合法
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "参数错误"
		ctx.JSON(r)
		return
	}

	// 获取路由参数中的用户名
	username := ctx.Params().Get("username")
	// 校验用户名和密码
	user, err := service.GetUserService().Login(username, body.Password)
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "登录失败, 用户名错误或密码错误"
		ctx.JSON(r)
		return
	}

	// 写入 session
	session := sessions.Get(ctx)
	session.Set("user", user.Name)

	r := &response.UserResponse{}
	r.Ok = true
	r.User = user
	ctx.JSON(r)
}

// @summary 		登出
// @description	 	退出登录
// @produce 		json
// @tags			user
// @router 			/api/logout [post]
// @success 		200	{object} response.SuccessResponse
func (uc *UserController) Logout(ctx iris.Context) {

	// 写入 session
	session := sessions.Get(ctx)
	session.Set("user", "")

	r := &response.SuccessResponse{}
	r.Ok = true
	ctx.JSON(r)
}
