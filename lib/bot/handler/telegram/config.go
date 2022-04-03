package tghandler

import _handler "github.com/artnoi43/fngobot/lib/bot/handler"

// Config for bot handlers. Mostly controls timing
type Config struct {
	Handler _handler.Config `mapstructure:"handler" json:"handler"`
	Client  ClientConfig    `mapstructure:"client" json:"client"`
}

type ClientConfig struct {
	TimeoutSeconds int      `mapstructure:"timeout_seconds" json:"timeoutSeconds"`
	Verbose        bool     `mapstructure:"verbose" json:"verbose"`
	BotTokens      []string `mapstructure:"bot_tokens" json:"botTokens"`
}
