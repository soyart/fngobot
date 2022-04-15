package bitkub

const URL = "https://bitkub.com/api/market/ticker/"

type fetcher struct{}

func New() interface{} {
	return &fetcher{}
}
