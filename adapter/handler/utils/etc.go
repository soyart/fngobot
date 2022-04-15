package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	TimeFormat string = "2006-01-02 15:04:05"
)

var (
	Printer = message.NewPrinter(language.English)
)
