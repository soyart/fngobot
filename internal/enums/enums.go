package enums

const (
	Bar  = "=============================="
	CONF = "$HOME/.config/fngobot/config.yml"
)

var BotMap = map[InputCommand]BotType{
	QuoteCommand:    QuoteBot,
	TrackCommand:    TrackBot,
	AlertCommand:    AlertBot,
	HelpCommand:     HelpBot,
	HandlersCommand: HandlersBot,
}
