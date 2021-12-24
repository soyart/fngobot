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
	HANDLERS
)

type Handler interface {
	UUID() string
	QuitChan() chan bool
	GetCmd() *parse.BotCommand
	send(string)
	Handle(int)
	HandleParsingError(parse.ParseError)
	SendQuote([]bot.Security)
	Track([]bot.Security, int, Config)
	PriceAlert(bot.Alert, Config)
}

type Handlers []Handler
var BotHandlers Handlers

type handler struct {
	Uuid string `json:"uuid"`
	Cmd  *parse.BotCommand `json:"command"`
	quit chan bool `json:"-"`
	conf Config `json:"-"`
	bot  *tb.Bot `json:"-"`
	msg  *tb.Message `json:"-"`
}
func (h *handler) UUID() string {
	return h.Uuid
}
func (h *handler) QuitChan() chan bool {
	return h.quit
}
func (h *handler) GetCmd() *parse.BotCommand {
	return h.Cmd
}

// Handle calls different methods on h based on its function parameter
func (h *handler) Handle(t int) {
	switch t {
	case QUOTEBOT:
		h.SendQuote(h.Cmd.Quote.Securities)
	case TRACKBOT:
		h.Track(h.Cmd.Track.Securities, h.Cmd.Track.TrackTimes, h.conf)
	case ALERTBOT:
		h.PriceAlert(h.Cmd.Alert, h.conf)
	case HANDLERS:
		h.SendHandlers()
	}
}

func (h *handler) send(s string) {
	h.bot.Send(h.msg.Sender, s)
}

func (h *handler) notifyStop() {
	log.Printf("[%s]: Received stop signal", h.Uuid)
	h.send(fmt.Sprintf("Stopping %s", h.Uuid))
}

// Stop stops a handler with matching UUID
func (h *Handlers) Stop(uuid string) (i int, ok bool) {
	for idx, handler := range *h {
		switch uuid {
		case handler.UUID():
			log.Printf("[%s]: Sending quit signal\n", handler.UUID())
			quit := handler.QuitChan()
			quit <- true
			log.Printf("[%s]: Sent quit signal\n", handler.UUID())
			i = idx
			ok = true
		}
	}
	return i, ok
}

// NewHandler returns a new handler
func NewHandler(b *tb.Bot, m *tb.Message, conf Config, cmd *parse.BotCommand) Handler {
	uuid := strings.Split(uuid.NewString(), "-")[0]
	quit := make(chan bool, 1)
	log.Printf("[%s]: %s (from %d)\n", uuid, m.Text, m.Sender.ID)
	h := &handler{
		Uuid: uuid,
		Cmd:  cmd,
		quit: quit,
		conf: conf,
		bot:  b,
		msg:  m,
	}
	BotHandlers = append(BotHandlers, h)
	return h
}