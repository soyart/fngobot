package handler

import (
	"log"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/parse"
)

type Handler interface {
	// These exported methods are called from other packages
	UUID() string
	QuitChan() chan struct{}
	Done()
	GetCmd() *parse.BotCommand
	Handle(enums.BotType)
	HandleParsingError(parse.ParseError)
	Quote([]bot.Security)
	Track([]bot.Security, int)
	PriceAlert(bot.Alert)
}

type Handlers []Handler

// Stop stops a handler with matching UUID,
// returning the index to the target handler
func (h Handlers) Stop(uuid string) (i int, ok bool) {
	for idx, handler := range h {
		switch uuid {
		case handler.UUID():
			log.Printf(
				"[%s]: sending quit signal\n",
				handler.UUID(),
			)
			quit := handler.QuitChan()
			quit <- struct{}{}
			log.Printf(
				"[%s]: sent quit signal\n",
				handler.UUID(),
			)
			i, ok = idx, true
		}
	}
	return i, ok
}
