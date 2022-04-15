package yahoocrypto

type fetcher struct{}

func New() interface{} {
	return &fetcher{}
}
