package enums

const (
	Bar  = "=============================="
	CONF = "$HOME/.config/fngobot/config.yml"
)

type (
	BotType   string
	Command   string
	QuoteType string
	Condition string
)

const (
	QuoteBot    BotType = "QUOTE"
	TrackBot    BotType = "TRACK"
	AlertBot    BotType = "ALERT"
	HelpBot     BotType = "HELP"
	HandlersBot BotType = "HANDLERSBOT"

	QuoteCommand    Command = "/quote"
	TrackCommand    Command = "/track"
	AlertCommand    Command = "/alert"
	HelpCommand     Command = "/help"
	HandlersCommand Command = "/handlers"

	Bid  QuoteType = "BID"
	Ask  QuoteType = "ASK"
	Last QuoteType = "LAST"

	Lt Condition = "LT"
	Gt Condition = "GT"
)

var BotMap = map[Command]BotType{
	QuoteCommand:    QuoteBot,
	TrackCommand:    TrackBot,
	AlertCommand:    AlertBot,
	HelpCommand:     HelpBot,
	HandlersCommand: HandlersBot,
}
