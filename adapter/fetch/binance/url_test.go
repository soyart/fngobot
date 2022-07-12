package binance

import (
	"testing"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
)

func TestPrepareUrl(t *testing.T) {
	info := &common.FetchInfo{
		URL:    BaseURL,
		Symbol: "ETHUSDT",
	}
	m := prepareURLs(info)
	if m["price"] != "https://api.binance.com/api/v3/ticker/price?symbol=ETHUSDT" {
		t.Error("price urls don't match")
	} else if m["depth"] != "https://api.binance.com/api/v3/ticker/bookTicker?symbol=ETHUSDT" {
		t.Error("depth urls don't match")
	}
}
