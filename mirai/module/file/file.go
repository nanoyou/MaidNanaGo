package file

import (
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/sirupsen/logrus"
)

var instance *File

func init() {
	instance = &File{}
	bot.RegisterModule(instance)
}

func GetInstance() *File {
	return instance
}

type File struct {
	bot.Module
}

func (v *File) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "md.file",
		Instance: instance,
	}
}

func (v *File) Init() {

}
func (v *File) PostInit() {

}

func (v *File) Serve(b *bot.Bot) {
	b.OnReceivedOfflineFile(func(q *client.QQClient, ofe *client.OfflineFileEvent) {
		logrus.WithField("file", ofe).Debug("接收到文件")
	})

	b.OnPrivateMessage(func(q *client.QQClient, pm *message.PrivateMessage) {
		for _, v := range pm.Elements {
			if v.Type() == message.Image {
				image := v.(*message.FriendImageElement)
				logrus.WithField("file", image).Debug("接收到图片")
				break
			}
		}
	})
}

func (v *File) Start(b *bot.Bot) {

}

func (v *File) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}
