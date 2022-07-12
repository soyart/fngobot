package binance

import (
	"fmt"
	"net/url"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
)

const BaseURL = "https://api.binance.com"

// Note to myself: tried escaping path and query,
// but that didn't work. Just don't bother.

// prepareURLs put string URLs into a map
func prepareURLs(info *common.FetchInfo) map[string]string {
	priceURL := preparePrice(info)
	depthURL := prepareDepth(info)

	return map[string]string{
		"price": priceURL,
		"depth": depthURL,
	}
}

// preparePrice just generate URL to
func preparePrice(info *common.FetchInfo) string {
	u, _ := url.Parse(info.URL)
	u.Path = "/api/v3/ticker/price"
	u.RawQuery = fmt.Sprintf(
		"symbol=%s",
		info.Symbol,
	)
	return u.String()
}
func prepareDepth(info *common.FetchInfo) string {
	u, _ := url.Parse(info.URL)
	u.Path = "/api/v3/ticker/bookTicker"
	u.RawQuery = fmt.Sprintf(
		"symbol=%s",
		info.Symbol,
	)
	return u.String()
}
