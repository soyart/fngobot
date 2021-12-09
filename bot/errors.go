package bot

import "errors"

var (
	errYahooZeroPrice    error = errors.New("fetched price is 0")
	errYahooCryptoBidAsk error = errors.New("bid/ask not supported")
	errSatangLastPrice   error = errors.New("last price not supported")
)
