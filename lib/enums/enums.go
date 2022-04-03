package enums

const (
	Bar  = "=============================="
	CONF = "$HOME/.config/fngobot/config.yml"
)

type (
	InputCommand string
	BotType      string
	QuoteType    string
	Condition    string
)

const (
	QuoteCommand    InputCommand = "/quote"
	TrackCommand    InputCommand = "/track"
	AlertCommand    InputCommand = "/alert"
	HelpCommand     InputCommand = "/help"
	HandlersCommand InputCommand = "/handlers"

	QuoteBot    BotType = "QUOTE"
	TrackBot    BotType = "TRACK"
	AlertBot    BotType = "ALERT"
	HelpBot     BotType = "HELP"
	HandlersBot BotType = "HANDLERSBOT"

	Bid  QuoteType = "BID"
	Ask  QuoteType = "ASK"
	Last QuoteType = "LAST"

	Lt Condition = "LT"
	Gt Condition = "GT"
)

var BotMap = map[InputCommand]BotType{
	QuoteCommand:    QuoteBot,
	TrackCommand:    TrackBot,
	AlertCommand:    AlertBot,
	HelpCommand:     HelpBot,
	HandlersCommand: HandlersBot,
}
