package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/nanoyou/MaidNanaGo/controller/request"
	"github.com/nanoyou/MaidNanaGo/controller/response"
	"github.com/nanoyou/MaidNanaGo/service"
)

type UserController struct{}

// @summary 		注册
// @description	 	注册账号, 需要先在 QQ 上私聊机器人发送 "验证码" 获取验证码
// @accept 			json
// @produce 		json
// @param			body body request.RegisterRequest true "注册信息"
// @tags			account user register
// @router 			/api/user [post]
// @success 		200	{object} response.RegisterSuccessResponse
// @failure 		200	{object} response.FailureResponse
func (uc *UserController) Register(ctx iris.Context) {
	var body request.RegisterRequest
	err := ctx.ReadJSON(&body)
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "参数错误"
		ctx.JSON(r)
		return
	}
	qq, err := service.GetUserService().GetQQByVerificationCode(body.VerificationCode)
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "验证码无效"
		ctx.JSON(r)
		return
	}

	user, err := service.GetUserService().Register(body.Username, body.Password, qq)
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "注册失败"
		ctx.JSON(r)
		return
	}
	r := &response.RegisterSuccessResponse{}
	r.Ok = true
	r.User = user
	ctx.JSON(r)

}
