package satang

type fetcher struct{}

func New() interface{} {
	return &fetcher{}
}

// Enum for parse() */
const (
	bid = iota
	ask
)
