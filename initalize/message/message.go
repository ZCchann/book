package message

import (
	"book/initalize/conf"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type lineBot struct {
	*linebot.Client
}

var l = new(lineBot)

func Line() *lineBot {
	return l
}

func (l *lineBot) InitLine() (err error) {
	bot, err := linebot.New(conf.Conf().LineBot.ChannelSecret, conf.Conf().LineBot.ChannelAccessToken)
	if err != nil {
		return
	}
	l.Client = bot
	return nil
}
