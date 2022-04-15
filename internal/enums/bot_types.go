package enums

type BotType string

const (
	QuoteBot    BotType = "QUOTE"
	TrackBot    BotType = "TRACK"
	AlertBot    BotType = "ALERT"
	HelpBot     BotType = "HELP"
	HandlersBot BotType = "HANDLERSBOT"
)

var (
	validBotTypes = []BotType{
		QuoteBot,
		TrackBot,
		AlertBot,
		HelpBot,
		HandlersBot,
	}
)
