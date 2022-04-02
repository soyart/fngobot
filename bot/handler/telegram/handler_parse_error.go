package tghandler

import (
	"strings"

	"github.com/artnoi43/fngobot/parse"
)

// HandleParsingError handles errors from package parse
func (h *handler) HandleParsingError(e parse.ParseError) {
	h.send(formString(e))
}

func formString(e parse.ParseError) string {
	signals := []string{
		"failed to parse command:",
		parse.ErrMsg[e],
	}
	return strings.Join(signals, "\n")
}
