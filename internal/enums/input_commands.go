package enums

type InputCommand string

const (
	QuoteCommand    InputCommand = "/quote"
	TrackCommand    InputCommand = "/track"
	AlertCommand    InputCommand = "/alert"
	HelpCommand     InputCommand = "/help"
	HandlersCommand InputCommand = "/handlers"
)

var (
	ValidInputCommands = []InputCommand{
		QuoteCommand,
		TrackCommand,
		AlertCommand,
		HelpCommand,
		HandlersCommand,
	}
)
