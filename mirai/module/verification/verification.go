package verification

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/nanoyou/MaidNanaGo/service"
)

var instance *Verification

func init() {
	instance = &Verification{}
	bot.RegisterModule(instance)
}

func GetInstance() *Verification {
	return instance
}

type Verification struct {
	bot.Module
}

func (v *Verification) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "md.verification",
		Instance: instance,
	}
}

func (v *Verification) Init() {

}
func (v *Verification) PostInit() {

}

func (v *Verification) Serve(b *bot.Bot) {
	b.OnPrivateMessageF(func(pm *message.PrivateMessage) bool {
		return strings.TrimSpace(pm.ToString()) == "验证码"
	}, func(q *client.QQClient, pm *message.PrivateMessage) {
		code, err := service.GetUserService().CreateVerificationCode(pm.Sender.Uin)
		if err != nil {
			b.SendPrivateMessage(pm.Sender.Uin, message.NewSendingMessage().Append(message.NewText("获取失败: "+err.Error())))
			return
		}
		b.SendPrivateMessage(pm.Sender.Uin, message.NewSendingMessage().Append(message.NewText(fmt.Sprintf("您的验证码是: %v", code))))
	})
}

func (v *Verification) Start(b *bot.Bot) {

}

func (v *Verification) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}
