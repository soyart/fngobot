package coinbase

type fetcher struct{}

func New() interface{} {
	return &fetcher{}
}

type info struct {
	symbol string
	url    string
}

type response struct {
	Data priceData `json:"data"`
}

type priceData struct {
	Amount string `json:"amount"`
}
