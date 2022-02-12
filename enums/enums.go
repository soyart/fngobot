package enums

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
