package tghandler

import "github.com/artnoi43/fngobot/parse"

// HandleParsingError handles errors from package parse
func (h *handler) HandleParsingError(e parse.ParseError) {
	h.send("Failed to parse command")
	h.send(parse.ErrMsg[e])
}
