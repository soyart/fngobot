package coinbase

import (
	"fmt"
	"net/url"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
)

const BaseURL = "https://api.coinbase.com"

// Note to myself: tried escaping path and query,
// but that didn't work. Just don't bother.

// prepareURLs put string URLs into a map
func prepareURLs(info *common.FetchInfo) map[string]string {
	priceURL := preparePrice(info)
	bidURL := prepareBid(info)
	askURL := prepareAsk(info)
	var m = make(map[string]string)
	m["price"] = priceURL
	m["bid"] = bidURL
	m["ask"] = askURL
	return m
}

// preparePrice just generate URL to
func preparePrice(info *common.FetchInfo) string {
	u, _ := url.Parse(info.URL)
	u.Path = fmt.Sprintf(
		"/v2/prices/%s-THB/spot",
		info.Symbol,
	)
	return u.String()
}
func prepareBid(info *common.FetchInfo) string {
	u, _ := url.Parse(info.URL)
	u.Path = fmt.Sprintf(
		"/v2/prices/%s-THB/buy",
		info.Symbol,
	)
	return u.String()
}
func prepareAsk(info *common.FetchInfo) string {
	u, _ := url.Parse(info.URL)
	u.Path = fmt.Sprintf(
		"/v2/prices/%s-THB/sell",
		info.Symbol,
	)
	return u.String()
}
