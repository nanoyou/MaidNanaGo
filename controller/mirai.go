package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/nanoyou/MaidNanaGo/controller/response"
	"github.com/nanoyou/MaidNanaGo/mirai/module/voice"
	"github.com/sirupsen/logrus"
)

type MiraiController struct{}

// @summary 		发送语音
// @description	 	发送私聊语音, 需要管理员权限
// @accept 			multipart/form-data
// @produce 		json
// @param 			qq path int true "QQ号"
// @param			voice formData file true "语音文件"
// @tags			admin, mirai
// @router 			/api/admin/mirai/voice/qq/{qq} [post]
// @success 		200	{object} response.SuccessResponse
// @failure 		200	{object} response.FailureResponse
func (mc *MiraiController) SendPrivateVoice(ctx iris.Context) {
	qq, err := ctx.Params().GetInt64("qq")
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "参数错误"
		ctx.JSON(r)
		return
	}
	f, fh, err := ctx.FormFile("voice")
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "文件读取失败"
		ctx.JSON(r)
		return
	}
	logrus.WithField("fileHeader", fh).Info("接收到音频文件")
	err = voice.GetInstance().SendPrivateVoice(qq, f)
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "发送失败"
		ctx.JSON(r)
		return
	}
	r := &response.SuccessResponse{}
	r.Ok = true
	r.SuccessMessage = "发送成功"
	ctx.JSON(r)

}
