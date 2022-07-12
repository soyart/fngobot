package usecase

import (
	"github.com/artnoi43/fngobot/internal/enums"
)

// GetSrcStr returns source in string based on s.Src
func (s *Security) GetSrcStr() string {
	return s.Src.String()
}

// GetCondStr returns condition in string based on a.Condition
func (a *Alert) GetCondStr() string {
	switch a.Condition {
	case enums.Lt:
		return "<="
	case enums.Gt:
		return ">="
	}
	return ">="
}

// GetQuoteTypeStr returns quote type in string based on a.QuoteType
func (a *Alert) GetQuoteTypeStr() string {
	switch a.QuoteType {
	case enums.Bid:
		return "BID"
	case enums.Ask:
		return "ASK"
	default:
		return "LAST"
	}
}
