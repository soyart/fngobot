package coinbase

type quote struct {
	last float64
	bid  float64
	ask  float64
}

func (q *quote) Last() (float64, error) {
	return q.last, nil
}
func (q *quote) Bid() (float64, error) {
	return q.bid, nil
}
func (q *quote) Ask() (float64, error) {
	return q.ask, nil
}
