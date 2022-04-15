package entity

// Quoter is returned by all Get functions
type Quoter interface {
	Last() (float64, error)
	Bid() (float64, error)
	Ask() (float64, error)
}
