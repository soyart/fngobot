package fetch

import (
	"errors"
)

var (
	NotFound error = errors.New("Ticker not found in JSON")
)
