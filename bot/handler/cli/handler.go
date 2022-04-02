package clihandler

import (
	"log"
	"time"

	_handler "github.com/artnoi43/fngobot/bot/handler"
	"github.com/artnoi43/fngobot/bot/utils"
	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/parse"
)

type handler struct {
	Uuid   string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	conf   *Config           `json:"-" yaml:"-"`
	Cmd    *parse.BotCommand `json:"command,omitempty" yaml:"command,omitempty"`
	Start  time.Time         `json:"start,omitempty" yaml:"start,omitempty"`
	IsDone bool              `json:"isDone" yaml:"isDone"`
	Quit   chan struct{}     `json:"-" yaml:"-"`
}

func (h *handler) UUID() string              { return h.Uuid }
func (h *handler) QuitChan() chan struct{}   { return h.Quit }
func (h *handler) GetCmd() *parse.BotCommand { return h.Cmd }
func (h *handler) Done()                     { h.IsDone = true }

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
		)
	case enums.ALERTBOT:
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
) _handler.Handler {
	return &handler{
		Start: time.Now(),
		Uuid:  utils.NewUUID(),
		Quit:  utils.NewQuit(),
		Cmd:   cmd,
		conf:  conf,
	}
}
