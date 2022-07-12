package common

// Quote struct for Bitkub
type Quote struct {
	Last float64
	Bid  float64
	Ask  float64
}

func (q *Quote) QuoteLast() (float64, error) {
	return q.Last, nil
}
func (q *Quote) QuoteBid() (float64, error) {
	return q.Bid, nil
}
func (q *Quote) QuoteAsk() (float64, error) {
	return q.Ask, nil
}
