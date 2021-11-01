package bot

import "github.com/artnoi43/fngobot/enums"

func (s *Security) GetSrcStr() string {
	switch s.Src {
	case enums.YahooCrypto:
		return "Crypto"
	case enums.Satang:
		return "Satang"
	case enums.Bitkub:
		return "Bitkub"
	default:
		return "Yahoo Finance"
	}
}

func (a *Alert) GetCondStr() string {
	switch a.Condition {
	case 0:
		return "<="
	default:
		return ">="
	}
}

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