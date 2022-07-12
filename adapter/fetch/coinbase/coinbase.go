package coinbase

type fetcher struct{}

func New() interface{} {
	return &fetcher{}
}

type response struct {
	Data priceData `json:"data"`
}

type priceData struct {
	Amount string `json:"amount"`
}
