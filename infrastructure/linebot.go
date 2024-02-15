package infrastructure

import (
	"github.com/JMjirapat/qrthrough-api/config"
	"github.com/line/line-bot-sdk-go/linebot"
)

var LineBot *linebot.Client

func InitLineBot() {
	cfg := config.Config

	var err error
	LineBot, err = linebot.New(cfg.ChannelSecret, cfg.ChannelToken)
	if err != nil {
		panic(err)
	}
}
