package coinbase

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
)

var (
	URLs = [3]string{
		BaseURL + "/v2/prices/BTC-THB/spot",
		BaseURL + "/v2/prices/BTC-THB/buy",
		BaseURL + "/v2/prices/XLM-THB/sell",
	}
)

func TestUnmarshal(t *testing.T) {
	fatal := func() {
		t.Fatalf("json unmarshal failed")
	}
	var wg sync.WaitGroup
	for _, u := range URLs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			body, err := common.Fetch(u)
			if err != nil {
				t.Errorf("fetch failed, %v", err.Error())
			}
			var r response
			if err := json.Unmarshal(body, &r); err != nil {
				fatal()
			}
			t.Log(r)
		}()
		wg.Wait()
	}
}
