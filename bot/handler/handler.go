package handler

import (
	"fmt"
	"log"
	"strings"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/parse"
	"github.com/google/uuid"
	tb "gopkg.in/tucnak/telebot.v3"
)

// Bot types - used in handler.Handle()
const (
	QUOTEBOT = iota
	TRACKBOT
	ALERTBOT
)

type Handler interface {
	Handle(int)
	HandleParsingError(int)
	SendQuote([]bot.Security)
	Track([]bot.Security, int, Config)
	PriceAlert(bot.Alert, Config)
}

var BotHandlers Handlers

// handler struct
type handler struct {
	uuid string
	quit chan bool
	conf Config
	cmd  *parse.BotCommand
	bot  *tb.Bot
	msg  *tb.Message
}

// NewHandler returns a new handler
func NewHandler(b *tb.Bot, m *tb.Message, conf Config, cmd *parse.BotCommand) Handler {
	uuid := strings.Split(uuid.NewString(), "-")[0]
	quit := make(chan bool, 1)
	log.Printf("[%s]: %s (from %d)\n", uuid, m.Text, m.Sender.ID)
	h := &handler{
		uuid: uuid,
		quit: quit,
		conf: conf,
		cmd:  cmd,
		bot:  b,
		msg:  m,
	}
	BotHandlers = append(BotHandlers, h)
	return h
}

// Handle calls different methods on h based on its function parameter
func (h *handler) Handle(t int) {
	switch t {
	case QUOTEBOT:
		h.SendQuote(h.cmd.Quote.Securities)
	case TRACKBOT:
		h.Track(h.cmd.Track.Securities, h.cmd.Track.TrackTimes, h.conf)
	case ALERTBOT:
		h.PriceAlert(h.cmd.Alert, h.conf)
	}
}

func (h *handler) send(s string) {
	h.bot.Send(h.msg.Sender, s)
}

func (h *handler) notifyStop() {
	log.Printf("[%s]: Received stop signal", h.uuid)
	h.send(fmt.Sprintf("Stopping %s", h.uuid))
}
