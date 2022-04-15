package binance

type info struct {
	symbol string
	url    string
}

type fetcher struct{}

func New() interface{} {
	return &fetcher{}
}
