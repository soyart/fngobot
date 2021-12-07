package bot

import "errors"

const (
	/* Error message: invalid ticker */
	apiError = "Ticker %s does not match any.\nFrom message:\n %s"
	etcError = "Error fetching price for %s\n\nFrom message:\n %s"
)

var (
	errYahoo error = errors.New("Fetched price is 0")
)
