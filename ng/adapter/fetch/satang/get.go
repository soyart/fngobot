package satang

import (
	"log"

	"github.com/artnoi43/fngobot/lib/fetch"
)

// Get fetches data from Satang JSON API,
// and parses the fetched JSON into Quote struct
func (f *fetcher) Get(tick string) (fetch.Quoter, error) {

	/* Documentation for Satang:
	 * https://docs.satangcorp.com/apis/public/orders */

	data, err := fetch.FetchMapStrInf("https://satangcorp.com/api/orderbook-tickers/")
	if err != nil {
		return nil, err
	}

	var q quote
	var found bool
	for key, val := range data {
		/* Filter ticker */
		switch key {
		case tick + "_THB":
			/* Inner keys and values */
			for k, v := range val.(map[string]interface{}) {
				switch k {
				case "bid":
					err := parse(&q, v, bid)
					if err != nil {
						return nil, err
					}
				case "ask":
					err := parse(&q, v, ask)
					if err != nil {
						return nil, err
					}
				}
			}
			found = true
		}
	}
	if !found {
		log.Printf("%s not found in Satang JSON", tick)
		return nil, fetch.ErrNotFound
	}
	return &q, nil
}
