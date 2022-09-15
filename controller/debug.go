package controller

import (
	"runtime/debug"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/kataras/iris/v12"
)

type DebugController struct{}

// @summary 		调试信息
// @description	 	获取 Maid Nana 调试信息
// @produce 		json
// @tags			about
// @router 			/api/about [get]
// @success 		200	{object} controller.DebugInfo
func (dc *DebugController) About(ctx iris.Context) {
	info := DebugInfo{}
	info.Version = "1.0.0-alpha"
	info.QQ.Online = bot.Instance.Online.Load()
	info.QQ.Account = bot.Instance.Uin
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		info.GoVersion = buildInfo.GoVersion
	}
	ctx.JSON(info)

}

type DebugInfo struct {
	Version   string
	GoVersion string
	QQ        struct {
		Account int64
		Online  bool `json:"online"`
	}
}
