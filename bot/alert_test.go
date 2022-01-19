package bot

import (
	"encoding/json"
	"sync"
	"testing"
	"time"

	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/fetch"
)

var (
	someSmallNum float64 = 0.69
	someBigNum   float64 = 200000000

	someSmallQuote quote = quote{
		bid:  someSmallNum,
		ask:  someSmallNum,
		last: someSmallNum,
	}
	someBigQuote quote = quote{
		bid:  someBigNum,
		ask:  someBigNum,
		last: someBigNum,
	}
	ethSatangGt0 = Alert{
		Security: Security{
			Tick: "ETH",
			Src:  enums.Satang,
		},
		QuoteType: enums.Bid,
		Condition: enums.Gt,
		Target:    0,
	}
	ethSatangLt5 = Alert{
		Security: Security{
			Tick: "ETH",
			Src:  enums.Satang,
		},
		QuoteType: enums.Bid,
		Condition: enums.Lt,
		Target:    5,
	}
	ethBitkubGt0 = Alert{
		Security: Security{
			Tick: "ETH",
			Src:  enums.Bitkub,
		},
		QuoteType: enums.Last,
		Condition: enums.Gt,
		Target:    0,
	}
	ethBitkubLt5 = Alert{
		Security: Security{
			Tick: "ETH",
			Src:  enums.Bitkub,
		},
		QuoteType: enums.Last,
		Condition: enums.Lt,
		Target:    5,
	}

	tests = map[Alert]map[fetch.Quoter]bool{
		// alert if: mk > 0
		// mk is now someBigNum
		// -> true
		ethSatangGt0: {
			someBigQuote: true,
		},
		// alert if: mk < 5
		// mk is now someSmallNum
		// -> true
		ethSatangLt5: {
			someSmallQuote: true,
		},
		// alert if: mk > 0
		// mk is now someBigNum
		// -> true
		ethBitkubGt0: {
			someBigQuote: true,
		},
		// alert if: mk < 5
		// mk is now someBigNum
		// -> false
		ethBitkubLt5: {
			someBigQuote: true,
		},
	}
)

func TestMatch(t *testing.T) {
	fail := func(c chan bool) {
		close(c)
		t.Fatal("results not expected")
	}
	var wg sync.WaitGroup
	for alert, resultMap := range tests {
		wg.Add(1)
		go func(a Alert, rm map[fetch.Quoter]bool) {
			defer wg.Done()
			for q, b := range rm {
				c := make(chan bool)
				e := make(chan error)
				go Match(&a, c, e, q)
				go func(e bool, mk fetch.Quoter) {
					j, _ := json.Marshal(a)
					js := string(j)
					for m := range c {
						t.Log(
							"alert", js,
							"market", mk,
							"expected", e,
							"actual", m,
						)
						if m != e {
							fail(c)
						}
					}
				}(b, q)
			}
		}(alert, resultMap)
	}
	wg.Wait()
	time.Sleep(2 * time.Second)
}

type quote struct {
	bid  float64
	ask  float64
	last float64
}

func (q quote) Bid() (float64, error) {
	return q.bid, nil
}
func (q quote) Ask() (float64, error) {
	return q.ask, nil
}
func (q quote) Last() (float64, error) {
	return q.last, nil
}
