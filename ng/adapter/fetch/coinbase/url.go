package coinbase

import (
	"fmt"
	"net/url"
)

const BaseURL = "https://api.coinbase.com"

// Note to myself: tried escaping path and query,
// but that didn't work. Just don't bother.

// prepareURLs put string URLs into a map
func prepareURLs(info info) map[string]string {
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
func preparePrice(info info) string {
	u, _ := url.Parse(info.url)
	u.Path = fmt.Sprintf(
		"/v2/prices/%s-THB/spot",
		info.symbol,
	)
	return u.String()
}
func prepareBid(info info) string {
	u, _ := url.Parse(info.url)
	u.Path = fmt.Sprintf(
		"/v2/prices/%s-THB/buy",
		info.symbol,
	)
	return u.String()
}
func prepareAsk(info info) string {
	u, _ := url.Parse(info.url)
	u.Path = fmt.Sprintf(
		"/v2/prices/%s-THB/sell",
		info.symbol,
	)
	return u.String()
}
