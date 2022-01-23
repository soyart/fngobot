package tghandler

import (
	"fmt"
	"log"
	"time"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/bot/utils"
	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/parse"
	tb "gopkg.in/tucnak/telebot.v3"
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
	Track([]bot.Security, int, Config)
	PriceAlert(bot.Alert, Config)
	// These unexported methods are called from this package
	send(string)
	yaml() string
	isRunning() bool
}

type Handlers []Handler

var (
	// SenderHandlers is a map of sender's ID and Handlers
	// so handlers are locally specific to the senders.
	SenderHandlers = make(map[int64]Handlers)
)

type handler struct {
	Uuid   string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Cmd    *parse.BotCommand `json:"command,omitempty" yaml:"command,omitempty"`
	Start  time.Time         `json:"start,omitempty" yaml:"start,omitempty"`
	Quit   chan struct{}     `json:"-" yaml:"-"`
	IsDone bool              `json:"-" yaml:"-"`
	Conf   Config            `json:"-" yaml:"-"`
	Bot    *tb.Bot           `json:"-" yaml:"-"`
	Msg    *tb.Message       `json:"-" yaml:"-"`
}

func (h *handler) UUID() string {
	return h.Uuid
}
func (h *handler) QuitChan() chan struct{} {
	return h.Quit
}
func (h *handler) Done() {
	h.IsDone = true
}
func (h *handler) isRunning() bool {
	return !h.IsDone
}
func (h *handler) GetCmd() *parse.BotCommand {
	return h.Cmd
}

// Handle calls different methods on h based on its function parameter
func (h *handler) Handle(t enums.BotType) {
	switch t {
	case enums.QUOTEBOT:
		h.Quote(
			h.GetCmd().Quote.Securities,
		)
	case enums.TRACKBOT:
		h.Track(
			h.GetCmd().Track.Securities,
			h.GetCmd().Track.TrackTimes,
			h.Conf,
		)
	case enums.ALERTBOT:
		h.PriceAlert(
			h.GetCmd().Alert,
			h.Conf,
		)
	case enums.HANDLERS:
		h.SendHandlers()
	}
}

// send sends given string to the handler's sender
func (h *handler) send(s string) {
	h.Bot.Send(h.Msg.Sender, s)
}

// notifyStop sends a Telegram message to sender to signal
// that the handler has received a quit signal.
func (h *handler) notifyStop() {
	log.Printf("[%s]: Received stop signal", h.Uuid)
	h.send(fmt.Sprintf("Stopping %s", h.Uuid))
}

// Stop stops a handler with matching UUID,
// returning the index to the target handler
func (h Handlers) Stop(uuid string) (i int, ok bool) {
	for idx, handler := range h {
		switch uuid {
		case handler.UUID():
			log.Printf(
				"[%s]: Sending quit signal\n",
				handler.UUID(),
			)
			quit := handler.QuitChan()
			quit <- struct{}{}
			log.Printf(
				"[%s]: Sent quit signal\n",
				handler.UUID(),
			)
			i, ok = idx, true
		}
	}
	return i, ok
}

// Remove removes a handler with matching sender and index
func Remove(senderId int64, idx int) {
	SenderHandlers[senderId] = append(
		SenderHandlers[senderId][:idx],
		SenderHandlers[senderId][idx+1:]...,
	)
}

// New returns a new handler and appends it to SenderHandlers
func New(
	b *tb.Bot,
	m *tb.Message,
	cmd *parse.BotCommand,
	conf Config,
) Handler {
	uuid := utils.NewUUID()
	// Log every new handler
	log.Printf(
		"[%s]: %s (from %d)\n",
		uuid,
		m.Text,
		m.Sender.ID,
	)
	h := &handler{
		Start:  time.Now(),
		Uuid:   uuid,
		Cmd:    cmd,
		Quit:   utils.NewQuit(),
		IsDone: false,
		Conf:   conf,
		Bot:    b,
		Msg:    m,
	}
	SenderHandlers[m.Sender.ID] = append(SenderHandlers[m.Sender.ID], h)
	return h
}
