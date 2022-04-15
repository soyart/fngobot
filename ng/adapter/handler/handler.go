package handler

import (
	"log"
	"time"

	"github.com/artnoi43/fngobot/ng/adapter/parse"
	"github.com/artnoi43/fngobot/ng/internal/enums"
	"github.com/artnoi43/fngobot/ng/usecase"
)

type Handler interface {
	// These exported methods are called from other packages
	UUID() string
	QuitChan() chan struct{}
	Done()
	IsRunning() bool
	GetCmd() *parse.BotCommand
	Handle(enums.BotType)
	HandleParsingError(parse.ParseError)
	Quote([]usecase.Security)
	Track([]usecase.Security, int)
	PriceAlert(usecase.Alert)
	StartTime() time.Time
}

type Handlers []Handler

// BaseHandler represents most boiler plate fields for any handler structs.
// Embed it in your new handler, or ignore it completely
type BaseHandler struct {
	Uuid   string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Cmd    *parse.BotCommand `json:"command,omitempty" yaml:"command,omitempty"`
	Start  time.Time         `json:"start,omitempty" yaml:"start,omitempty"`
	IsDone bool              `json:"-" yaml:"-"`
	Quit   chan struct{}     `json:"-" yaml:"-"`
}

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

// These types are only for marshaling YAML
type PrettyAlert struct {
	Security  usecase.Security `yaml:"Security,omitempty"`
	Condition string           `yaml:"Condition,omitempty"`
	Target    float64          `yaml:"Target,omitempty"`
}
type PrettyHandler struct {
	Uuid  string             `yaml:"UUID,omitempty"`
	Start string             `yaml:"Start,omitempty"`
	Quote []usecase.Security `yaml:"Quote,omitempty"`
	Track []usecase.Security `yaml:"Track,omitempty"`
	Alert PrettyAlert        `yaml:"Alert,omitempty"`
}
