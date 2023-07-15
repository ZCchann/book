package line

import (
	"book/initalize/conf"
	"book/initalize/message"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
)

func SendMessage(text string) {
	if !conf.Conf().LineBot.State {
		return
	}
	_, err := message.Line().PushMessage(conf.Conf().LineBot.GroupID, linebot.NewTextMessage(text)).Do()
	if err != nil {
		log.Println(err)
		return
	}
}
