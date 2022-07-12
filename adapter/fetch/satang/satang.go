package satang

type fetcher struct{}

func New() interface{} {
	return &fetcher{}
}

type satangQuote struct {
	Price string `json:"price"`
}

type satangResponse map[string][]satangQuote
