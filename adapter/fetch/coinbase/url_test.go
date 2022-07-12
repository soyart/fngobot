package coinbase

import (
	"testing"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
)

func TestPrepareUrl(t *testing.T) {
	info := &common.FetchInfo{
		URL:    BaseURL,
		Symbol: "CAKE",
	}
	m := prepareURLs(info)
	if actual := m["price"]; actual != "https://api.coinbase.com/v2/prices/CAKE-THB/spot" {
		t.Logf(
			"Expected https://api.coinbase.com/v2/prices/CAKE-THB/spot, got %v",
			m["price"],
		)
		t.Error("price urls don't match")
	}
	if actual := m["bid"]; actual != "https://api.coinbase.com/v2/prices/CAKE-THB/buy" {
		t.Logf(
			"Expected https://api.coinbase.com/v2/prices/CAKE-THB/buy, got %v",
			m["bid"],
		)
		t.Error("bid urls don't match")
	}
	if actual := m["ask"]; actual != "https://api.coinbase.com/v2/prices/CAKE-THB/sell" {
		t.Logf(
			"Expected https://api.coinbase.com/v2/prices/CAKE-THB/buy, got %v",
			m["ask"],
		)
		t.Error("ask urls don't match")
	}
}
