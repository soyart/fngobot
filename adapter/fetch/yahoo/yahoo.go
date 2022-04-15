package yahoo

type fetcher struct{}

func New() interface{} {
	return &fetcher{}
}
