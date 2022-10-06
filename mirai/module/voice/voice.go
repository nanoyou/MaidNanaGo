package voice

import (
	"errors"
	"io"
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Mrs4s/MiraiGo/message"
)

var instance *Voice

func init() {
	instance = &Voice{}
	bot.RegisterModule(instance)
}

func GetInstance() *Voice {
	return instance
}

type Voice struct {
	bot.Module
	b *bot.Bot
}

func (v *Voice) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "md.voice",
		Instance: instance,
	}
}

func (v *Voice) Init() {

}
func (v *Voice) PostInit() {

}

func (v *Voice) Serve(b *bot.Bot) {
}

func (v *Voice) Start(b *bot.Bot) {
	v.b = b
}

func (v *Voice) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}

func (v *Voice) SendPrivateVoice(qq int64, voice io.ReadSeeker) error {
	element, err := bot.Instance.UploadPrivatePtt(qq, voice)
	if err != nil {
		return err
	}
	if v.b.SendPrivateMessage(qq, message.NewSendingMessage().Append(element)) == nil {
		return errors.New("发送失败")
	}
	return nil
}
