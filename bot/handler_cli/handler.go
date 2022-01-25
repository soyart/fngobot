package clihandler

import (
	"fmt"
	"time"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/bot/utils"
	"github.com/artnoi43/fngobot/config"
	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/parse"
)

type Handler interface {
	UUID() string
	Done()
	QuitChan() chan struct{}
	isRunning() bool
	GetCmd() *parse.BotCommand
	Handle(enums.BotType)
	HandleParsingError(parse.ParseError)
	Quote([]bot.Security)
	Track([]bot.Security, int, *config.Config)
	PriceAlert(bot.Alert, *config.Config)
}

type Handlers []Handler

type handler struct {
	Uuid   string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Cmd    *parse.BotCommand `json:"command,omitempty" yaml:"command,omitempty"`
	Start  time.Time         `json:"start,omitempty" yaml:"start,omitempty"`
	IsDone bool              `json:"isDone" yaml:"isDone"`
	Quit   chan struct{}     `json:"-" yaml:"-"`
	conf   *config.Config    `json:"-" yaml:"-"`
}

func (h *handler) UUID() string              { return h.Uuid }
func (h *handler) QuitChan() chan struct{}   { return h.Quit }
func (h *handler) isRunning() bool           { return !h.IsDone }
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
			h.conf,
		)
	case enums.ALERTBOT:
		h.PriceAlert(
			h.GetCmd().Alert,
			h.conf,
		)
	}
}

func (h *handler) HandleParsingError(e parse.ParseError) {
	fmt.Printf(
		"[error] %s\n",
		parse.ErrMsg[e],
	)
}

func (h *handler) notifyStop() {
	fmt.Printf("Stopping %s\n", h.UUID())
}

func New(
	cmd *parse.BotCommand,
	conf *config.Config,
) Handler {
	return &handler{
		Start: time.Now(),
		Uuid:  utils.NewUUID(),
		Quit:  utils.NewQuit(),
		Cmd:   cmd,
		conf:  conf,
	}
}
