package bot

import "github.com/artnoi43/fngobot/enums"

// GetSrcStr returns source in string based on s.Src
func (s *Security) GetSrcStr() string {
	switch s.Src {
	case enums.YahooCrypto:
		return "Crypto"
	case enums.Satang:
		return "Satang"
	case enums.Bitkub:
		return "Bitkub"
	case enums.Binance:
		return "Binance"
	default:
		return "Yahoo Finance"
	}
}

// GetCondStr returns condition in string based on a.Condition
func (a *Alert) GetCondStr() string {
	switch a.Condition {
	case enums.Lt:
		return "<="
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
