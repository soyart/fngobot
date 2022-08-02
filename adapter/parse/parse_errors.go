package parse

import "errors"

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

var ErrMsgs = map[ParseError]error{
	ErrParseInt:             errors.New("failed to parse integer. For /track, you must supply the tracking rounds (last argumemt) in integer"),
	ErrParseFloat:           errors.New("failed to parse float. For /alert, you must supply the target price (last argument) in floating point number"),
	ErrInvalidSign:          errors.New("failed to parse comparison sign. For /alert, you may only supply '>' or '<' before the target price as comparison condition"),
	ErrInvalidBidAskSwitch:  errors.New("failed to parse bid/ask switch. For /alert with bid/ask price, you may only supply 'bid' or 'ask' before the comparison sign"),
	ErrInvalidQuoteTypeBid:  errors.New("source does not support bid price. Omit the bid keyword"),
	ErrInvalidQuoteTypeAsk:  errors.New("source does not support ask price. Omit the ask keyword"),
	ErrInvalidQuoteTypeLast: errors.New("source does not support last price. Add either 'bid' or 'ask' keyword before the comparison sign"),
}
