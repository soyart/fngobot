package binance

type fetcher struct{}

func New() interface{} {
	return &fetcher{}
}
