package bot

import "errors"

const (
	/* Alert times */
	TIMES = 5
	/* Error message: invalid ticker */
	apiError = "Ticker %s does not match any.\nFrom message:\n %s"
	etcError = "Error fetching price for %s\n\nFrom message:\n %s"
)

var (
	yahooError error = errors.New("Fetched price is 0")
)