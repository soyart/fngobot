package entity

// Quoter is returned by all Get functions
type Quoter interface {
	QuoteLast() (float64, error)
	QuoteBid() (float64, error)
	QuoteAsk() (float64, error)
}
