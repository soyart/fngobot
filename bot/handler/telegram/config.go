package tghandler

import _handler "github.com/artnoi43/fngobot/bot/handler"

// Config for bot handlers. Mostly controls timing
type Config struct {
	Handler   _handler.Config `mapstructure:"handler" json:"handler"`
	BotTokens []string        `mapstructure:"bot_tokens" json:"botTokens"`
}
