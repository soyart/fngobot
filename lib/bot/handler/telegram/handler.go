package tghandler

import (
	"fmt"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v3"

	_handler "github.com/artnoi43/fngobot/lib/bot/handler"
	"github.com/artnoi43/fngobot/lib/bot/handler/utils"
	"github.com/artnoi43/fngobot/lib/enums"
	"github.com/artnoi43/fngobot/lib/parse"
)

var (
	// SenderHandlers is a map of sender's ID and Handlers
	// so handlers are locally specific to the senders.
	SenderHandlers = make(map[int64]_handler.Handlers)
)

// handler implements _handler.Handler, and has _handler.BaseHandler embedded
type handler struct {
	*_handler.BaseHandler
	conf Config     `json:"-" yaml:"-"`
	c    tb.Context `json:"-" yaml:"-"`
	bot  *tb.Bot    `json:"-" yaml:"-"`
}

func (h *handler) UUID() string              { return h.Uuid }
func (h *handler) QuitChan() chan struct{}   { return h.Quit }
func (h *handler) isRunning() bool           { return !h.IsDone }
func (h *handler) GetCmd() *parse.BotCommand { return h.Cmd }
func (h *handler) Done()                     { h.IsDone = true }

// Handle calls different methods on h based on its function parameter
func (h *handler) Handle(t enums.BotType) {
	switch t {
	case enums.QuoteBot:
		h.Quote(
			h.GetCmd().Quote.Securities,
		)
	case enums.TrackBot:
		h.Track(
			h.GetCmd().Track.Securities,
			h.GetCmd().Track.TrackTimes,
		)
	case enums.AlertBot:
		h.PriceAlert(
			h.GetCmd().Alert,
		)
	case enums.HandlersBot:
		if err := h.SendHandlers(); err != nil {
			log.Println("failed to send handlers:", err.Error())
		}
	}
}

// send sends given string to the handler's sender
// Now fngobot uses reply()
// func (h *handler) send(s string) {
// 	sender := h.c.Message().Sender
// 	if _, err := h.Bot.Send(sender, s); err != nil {
// 		log.Printf("[%s] failed to send message\n", h.Uuid)
// 	}
// }

func (h *handler) reply(s string) {
	chatMsg := h.c.Message()
	if _, err := h.bot.Reply(chatMsg, s); err != nil {
		log.Printf("[%s] failed to reply message\n", h.Uuid)
	}
}

// notifyStop sends a Telegram message to sender to signal
// that the handler has received a quit signal.
func (h *handler) notifyStop() {
	log.Printf("[%s]: received stop signal", h.Uuid)
	h.reply(fmt.Sprintf("Stopping %s", h.Uuid))
}

// Remove removes a handler with matching sender and index
func Remove(senderId int64, idx int) {
	SenderHandlers[senderId] = append(
		SenderHandlers[senderId][:idx],
		SenderHandlers[senderId][idx+1:]...,
	)
}

// New returns a new handler (Telegram) and appends it to SenderHandlers
func New(
	b *tb.Bot,
	c tb.Context,
	cmd *parse.BotCommand,
	conf Config,
) _handler.Handler {
	uuid := utils.NewUUID()
	m := c.Message()
	// Log every new handler
	log.Printf(
		"[%s]: %s (from %d)\n",
		uuid, m.Text, m.Sender.ID,
	)
	h := &handler{
		BaseHandler: &_handler.BaseHandler{
			Start:  time.Now(),
			Uuid:   uuid,
			Cmd:    cmd,
			IsDone: false,
			Quit:   utils.NewQuit(),
		},
		conf: conf,
		bot:  b,
		c:    c,
	}
	SenderHandlers[m.Sender.ID] = append(SenderHandlers[m.Sender.ID], h)
	return h
}
