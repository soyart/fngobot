package handler

import "github.com/artnoi43/fngobot/parse"

func (h *Handler) HandleParsingError(parseError int) {
	h.send("Failed to parse command")
	switch parseError {
	case parse.ErrParseInt:
		h.send("Failed to parse integer\nFor /track, you must supply the tracking rounds (last argumemt) in integer")
	case parse.ErrParseFloat:
		h.send("Failed to parse float\nFor /alert, you must supply the target price (last argument) in floating point number")
	case parse.ErrInvalidSign:
		h.send("Failed to parse comparison sign\nFor /alert, you may only supply '>' or '<' before the target price as comparison condition")
	case parse.ErrInvalidBidAskSwitch:
		h.send("Failed to parse bid/ask switch\nFor /alert with bid/ask price, you may only supply 'bid' or 'ask' before the comparison sign")
	case parse.ErrInvalidQuoteTypeBid:
		h.send("This source does not support bid price\nPlease omit the bid keyword")
	case parse.ErrInvalidQuoteTypeAsk:
		h.send("This source does not support ask price\nPlease omit the ask keyword\n")
	case parse.ErrInvalidQuoteTypeLast:
		h.send("This source does not support last price\nPlease add bid/ask switch before the comparison sign")
	default:

	}
}
