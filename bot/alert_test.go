package bot

import (
	"encoding/json"
	"sync"
	"testing"
	"time"

	"github.com/artnoi43/fngobot/enums"
)

var (
	someBigNum float64 = 20000000
	// (Should be true) Satang ETH > 0
	ethSatangGt0 = Alert{
		Security: Security{
			Tick: "ETH",
			Src:  enums.Satang,
		},
		QuoteType: enums.Bid,
		Condition: enums.Gt,
		Target:    0,
	}
	// (Should be false) Satang ETH > someBigNum
	ethSatangLtB = Alert{
		Security: Security{
			Tick: "ETH",
			Src:  enums.Satang,
		},
		QuoteType: enums.Bid,
		Condition: enums.Lt,
		Target:    someBigNum,
	}
	// (Should be true) Bitkub ETH > 0
	ethBitkubGt0 = Alert{
		Security: Security{
			Tick: "ETH",
			Src:  enums.Bitkub,
		},
		QuoteType: enums.Last,
		Condition: enums.Gt,
		Target:    0,
	}
	// (Should be false) Bitkub ETH > someBigNum
	ethBitkubLtB = Alert{
		Security: Security{
			Tick: "ETH",
			Src:  enums.Bitkub,
		},
		QuoteType: enums.Last,
		Condition: enums.Lt,
		Target:    someBigNum,
	}

	tests = map[Alert]bool{
		// Note: Match() will send false to a channel
		// only when there's an error when getting quotes
		ethSatangGt0: true,
		ethSatangLtB: true,
		ethBitkubGt0: true,
		ethBitkubLtB: true,
	}
)

func TestMatch(t *testing.T) {
	fail := func() {
		t.Fatal("results not expected")
	}
	var wg sync.WaitGroup
	for alert, expectedResult := range tests {
		wg.Add(1)
		go func(a Alert, b bool) {
			go func(expected bool) {
				c := make(chan bool)
				defer wg.Done()
				go Match(&a, c)
				go func(e bool) {
					for m := range c {
						// Marshal our alert into JSON
						j, _ := json.Marshal(a)
						js := string(j)
						if m != e {
							t.Log(
								"alert", js,
								"expected", e,
								"actual", m,
							)
							fail()
						}
					}
				}(expected)
			}(b)
		}(alert, expectedResult)
	}
	time.Sleep(time.Second)
	wg.Wait()
}
