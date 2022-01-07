package handler

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	timeFormat string = "2006-01-02 15:04:05"
)

var (
	printer = message.NewPrinter(language.English)
)
