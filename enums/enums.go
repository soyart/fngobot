package enums

import "fmt"

const CONF = "$HOME/.config/fngobot/config.yml"

// Bot types - used in handler.Handle()
type BotType string

const (
	QUOTEBOT BotType = "Quote"
	TRACKBOT BotType = "Track"
	ALERTBOT BotType = "Alert"
	HELPBOT  BotType = "Help"
	HANDLERS BotType = "Handlers"
)

type Src string

// quote sources - when adding new sources,
// also add them to validSrc below
const (
	Yahoo       Src = "Yahoo"
	YahooCrypto Src = "YahooCrypto"
	Satang      Src = "Satang"
	Bitkub      Src = "Bitkub"
	Binance     Src = "Binance"
	Coinbase    Src = "Coinbase"
)

var validSrc = [6]Src{
	Yahoo,
	YahooCrypto,
	Satang,
	Bitkub,
	Binance,
	Coinbase,
}

var ErrInvalidSrc = fmt.Errorf("invalid source")

func (s Src) IsValid() bool {
	for _, valid := range validSrc {
		if s == valid {
			return true
		}
	}
	return false
}

// quote types
type QuoteType string

const (
	Bid  QuoteType = "bid"
	Ask  QuoteType = "ask"
	Last QuoteType = "last"
)

// alert condition
type Condition string

const (
	Lt Condition = "lt"
	Gt Condition = "gt"
)

// For CLI
type Command string

const (
	QuoteCommand    Command = "/quote"
	TrackCommand    Command = "/track"
	AlertCommand    Command = "/alert"
	HelpCommand     Command = "/help"
	HandlersCommand Command = "/handlers"
)

var BotMap = map[Command]BotType{
	QuoteCommand:    QUOTEBOT,
	TrackCommand:    TRACKBOT,
	AlertCommand:    ALERTBOT,
	HelpCommand:     HELPBOT,
	HandlersCommand: HANDLERS,
}
