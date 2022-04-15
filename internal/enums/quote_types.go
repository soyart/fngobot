package enums

type QuoteType string

const (
	Bid  QuoteType = "BID"
	Ask  QuoteType = "ASK"
	Last QuoteType = "LAST"
)

var (
	validQuoteTypes = []QuoteType{
		Last,
		Bid,
		Ask,
	}
)
