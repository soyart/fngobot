package bitkub

const (
	BaseURL      = "https://bitkub.com/api/market/ticker/"
	TickerPrefix = "THB_"
)

type tokenInnerInfo struct {
	Last float64 `json:"last"`
	Ask  float64 `json:"lowestAsk"`
	Bid  float64 `json:"highestBid"`
}

type bitkubResponse struct {
	DataPart map[string]tokenInnerInfo `json:"data"`
}

type fetcher struct{}

func New() interface{} {
	return &fetcher{}
}
