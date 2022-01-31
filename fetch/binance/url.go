package binance

import (
	"fmt"
	"net/url"
)

// Note to myself: tried escaping path and query,
// but that didn't work. Just don't bother.

// prepareURLs put string URLs into a map
func prepareURLs(info info) map[string]string {
	priceURL := preparePrice(info)
	depthURL := prepareDepth(info)

	return map[string]string{
		"price": priceURL,
		"depth": depthURL,
	}
}

// preparePrice just generate URL to
func preparePrice(info info) string {
	u, _ := url.Parse(info.url)
	u.Path = "/api/v3/ticker/price"
	u.RawQuery = fmt.Sprintf(
		"symbol=%s",
		info.symbol,
	)
	return u.String()
}
func prepareDepth(info info) string {
	u, _ := url.Parse(info.url)
	u.Path = "/api/v3/ticker/bookTicker"
	u.RawQuery = fmt.Sprintf(
		"symbol=%s",
		info.symbol,
	)
	return u.String()
}
