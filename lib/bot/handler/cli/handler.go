package clihandler

import (
	"log"
	"time"

	_handler "github.com/artnoi43/fngobot/lib/bot/handler"
	"github.com/artnoi43/fngobot/lib/bot/handler/utils"
	"github.com/artnoi43/fngobot/lib/enums"
	"github.com/artnoi43/fngobot/lib/parse"
)

type handler struct {
	*_handler.BaseHandler
	conf *Config       `json:"-" yaml:"-"`
	done chan struct{} `json:"-" yaml:"-"`
}

func (h *handler) UUID() string              { return h.Uuid }
func (h *handler) QuitChan() chan struct{}   { return h.Quit }
func (h *handler) GetCmd() *parse.BotCommand { return h.Cmd }
func (h *handler) Done()                     { h.IsDone = true }

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

func (h *handler) HandleParsingError(e parse.ParseError) {
	log.Printf(
		"[error] %s\n",
		parse.ErrMsg[e],
	)
}

func (h *handler) notifyStop() {
	log.Printf("Stopping %s\n", h.UUID())
}

func New(
	cmd *parse.BotCommand,
	conf *Config,
	done chan struct{},
) _handler.Handler {
	return &handler{
		BaseHandler: &_handler.BaseHandler{
			Start: time.Now(),
			Uuid:  utils.NewUUID(),
			Cmd:   cmd,
			Quit:  utils.NewQuit(),
		},
		conf: conf,
		done: done,
	}
}
