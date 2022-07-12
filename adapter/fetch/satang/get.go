package satang

import (
	"log"
	"strconv"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
	"github.com/artnoi43/fngobot/usecase"
)

const BaseURL = "https://satangcorp.com/api/orderbook-tickers/"

// Get fetches data from Satang JSON API,
// and parses the fetched JSON into Quote struct
func (f *fetcher) Get(tick string) (usecase.Quoter, error) {
	/* Documentation for Satang:
	 * https://docs.satangcorp.com/apis/public/orders */
	data, err := common.FetchMapStrInf(BaseURL)
	if err != nil {
		return nil, err
	}

	var q common.Quote
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
		log.Printf("%s not found in Satang JSON\n", tick)
		return nil, common.ErrNotFound
	}
	return &q, nil
}

func parse(q *common.Quote, val interface{}, bidAsk int) error {
	for k, v := range val.(map[string]interface{}) {
		switch k {
		case "price":
			price, err := strconv.ParseFloat(v.(string), 64)
			if err != nil {
				return err
			}
			switch bidAsk {
			case bid:
				q.Bid = price
			case ask:
				q.Ask = price
			}
		}
	}
	return nil
}
