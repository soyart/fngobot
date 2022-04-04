package parse

type ParseError int

const (
	// These errors are non-zero
	NoErr ParseError = iota
	ErrParseInt
	ErrParseFloat
	ErrInvalidSign
	ErrInvalidBidAskSwitch
	ErrInvalidQuoteTypeBid
	ErrInvalidQuoteTypeAsk
	ErrInvalidQuoteTypeLast
)

var ErrMsg = map[ParseError]string{
	ErrParseInt:            "Failed to parse integer. For /track, you must supply the tracking rounds (last argumemt) in integer",
	ErrParseFloat:          "Failed to parse float. For /alert, you must supply the target price (last argument) in floating point number",
	ErrInvalidSign:         "Failed to parse comparison sign. For /alert, you may only supply '>' or '<' before the target price as comparison condition",
	ErrInvalidBidAskSwitch: "Failed to parse bid/ask switch. For /alert with bid/ask price, you may only supply 'bid' or 'ask' before the comparison sign",
	ErrInvalidQuoteTypeBid: "This source does not support bid price. Omit the bid keyword",
	ErrInvalidQuoteTypeAsk: "This source does not support ask price. Omit the ask keyword",
}
