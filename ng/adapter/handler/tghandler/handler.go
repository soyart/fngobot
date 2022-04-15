package tghandler

import (
	"fmt"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v3"

	_handler "github.com/artnoi43/fngobot/ng/adapter/handler"
	"github.com/artnoi43/fngobot/ng/adapter/handler/utils"
	"github.com/artnoi43/fngobot/ng/adapter/parse"
	"github.com/artnoi43/fngobot/ng/internal/enums"
)

// handler implements _handler.Handler, and has _handler.BaseHandler embedded
type handler struct {
	*_handler.BaseHandler
	conf Config     `json:"-" yaml:"-"`
	c    tb.Context `json:"-" yaml:"-"`
	bot  *tb.Bot    `json:"-" yaml:"-"`
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
	return h
}

func (h *handler) UUID() string              { return h.Uuid }
func (h *handler) QuitChan() chan struct{}   { return h.Quit }
func (h *handler) isRunning() bool           { return !h.IsDone }
func (h *handler) GetCmd() *parse.BotCommand { return h.Cmd }
func (h *handler) Done()                     { h.IsDone = true }
func (h *handler) IsRunning() bool           { return !h.IsDone }

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
	}
}

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

func (h *handler) StartTime() time.Time {
	return h.Start
}
