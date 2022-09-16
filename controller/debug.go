package controller

import (
	"runtime/debug"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/kataras/iris/v12"
	"github.com/nanoyou/MaidNanaGo/controller/response"
)

type DebugController struct{}

// @summary 		调试信息
// @description	 	获取 Maid Nana 调试信息
// @produce 		json
// @tags			about
// @router 			/api/about [get]
// @success 		200	{object} response.DebugInfo
func (dc *DebugController) About(ctx iris.Context) {
	info := response.DebugInfo{}
	info.Version = "1.0.0-alpha"
	info.QQ.Online = bot.Instance.Online.Load()
	info.QQ.Account = bot.Instance.Uin
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		info.GoVersion = buildInfo.GoVersion
	}
	ctx.JSON(info)

}
