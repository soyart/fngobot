package usecase

import (
	"encoding/json"
	"sync"
	"testing"
	"time"

	"github.com/artnoi43/fngobot/entity"
	"github.com/artnoi43/fngobot/internal/enums"
)

var (
	someSmallNum float64 = 0.69
	someBigNum   float64 = 200000000

	someSmallQuote = quote{
		bid:  someSmallNum,
		ask:  someSmallNum,
		last: someSmallNum,
	}
	someBigQuote = quote{
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

	tests = map[Alert]map[entity.Quoter]bool{
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
		// alert if: mk < 5
		// mk is now someBigNum
		// -> false
		ethBitkubLt5: {
			someBigQuote: false,
		},
		// alert if: mk > 0
		// mk is now someBigNum
		// -> true
		ethBitkubGt0: {
			someBigQuote: true,
		},
		// alert if: mk < 5
		// mk is now someSmallNum
		// -> true
		ethBitkubLt5: {
			someSmallQuote: true,
		},
		// alert if: mk < 5
		// mk is now someBigNum
		// -> false
		ethBitkubLt5: {
			someBigQuote: false,
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
		go func(a Alert, rm map[entity.Quoter]bool) {
			defer wg.Done()
			for q, b := range rm {
				c := make(chan bool)
				e := make(chan error)
				go Match(&a, c, e, q)
				go func(e bool, mk entity.Quoter) {
					j, _ := json.Marshal(a)
					js := string(j)
					for m := range c {
						if m != e {
							t.Log(
								"alert", js,
								"market", mk,
								"expected", e,
								"actual", m,
							)
							fail(c)
						}
					}
				}(b, q)
			}
		}(alert, resultMap)
	}
	time.Sleep(20 * time.Millisecond)
	wg.Wait()
}

type quote struct {
	bid  float64
	ask  float64
	last float64
}

func (q quote) QuoteBid() (float64, error) {
	return q.bid, nil
}
func (q quote) QuoteAsk() (float64, error) {
	return q.ask, nil
}
func (q quote) QuoteLast() (float64, error) {
	return q.last, nil
}
