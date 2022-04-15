package tgdriver

import (
	"log"

	tb "gopkg.in/tucnak/telebot.v3"

	"github.com/artnoi43/fngobot/adapter/handler/tghandler"
)

type TgDriver interface {
	InitAndStartBot() error
}

type tgDriver struct {
	bot           *tb.Bot
	token         string
	handlerConfig *tghandler.Config
}

func New(b *tb.Bot, token string, conf *tghandler.Config) TgDriver {
	if b == nil {
		log.Fatalln("nil bot for", token)
	}
	return &tgDriver{
		bot:           b,
		token:         token,
		handlerConfig: conf,
	}
}
