package handler

import "github.com/artnoi43/fngobot/parse"

func (h *Handler) HandleParsingError(parseError int) {
	h.bot.Send(h.msg.Sender, "Failed to parse command")
	switch parseError {
	case parse.ErrParseInt:
		h.bot.Send(h.msg.Sender, "Failed to parse integer\nFor /track, you must supply the tracking rounds (last argumemt) in integer")
	case parse.ErrParseFloat:
		h.bot.Send(h.msg.Sender, "Failed to parse float\nFor /alert, you must supply the target price (last argument) in floating point number")
	case parse.ErrInvalidSign:
		h.bot.Send(h.msg.Sender, "Failed to parse comparison sign\nFor /alert, you may only supply '>' or '<' before the target price as comparison condition")
	case parse.ErrInvalidBidAskSwitch:
		h.bot.Send(h.msg.Sender, "Failed to parse bid/ask switch\nFor /alert with bid/ask price, you may only supply 'bid' or 'ask' before the comparison sign")
	case parse.ErrInvalidQuoteTypeBid:
		h.bot.Send(h.msg.Sender, "This source does not support bid price\nPlease omit the bid keyword")
	case parse.ErrInvalidQuoteTypeAsk:
		h.bot.Send(h.msg.Sender, "This source does not support ask price\nPlease omit the ask keyword\n")
	case parse.ErrInvalidQuoteTypeLast:
		h.bot.Send(h.msg.Sender, "This source does not support last price\nPlease add bid/ask switch before the comparison sign")
	default:

	}
}
