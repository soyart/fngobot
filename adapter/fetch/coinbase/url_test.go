package coinbase

import (
	"testing"
)

func TestPrepareUrl(t *testing.T) {
	info := info{
		symbol: "CAKE",
		url:    BaseURL,
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
