package clihandler

import (
	_handler "github.com/artnoi43/fngobot/adapter/handler"
)

type Config struct {
	Handler _handler.Config `mapstructure:"handler" json:"handler"`
}
