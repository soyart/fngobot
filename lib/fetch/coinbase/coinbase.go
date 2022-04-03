package coinbase

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/pkg/errors"

	"github.com/artnoi43/fngobot/lib/fetch"
)

const BaseURL = "https://api.coinbase.com"

type info struct {
	symbol string
	url    string
}

type response struct {
	Data priceData `json:"data"`
}

type priceData struct {
	Amount string `json:"amount"`
}

func Get(tick string) (fetch.Quoter, error) {
	urls := prepareURLs(info{
		symbol: tick,
		url:    BaseURL,
	})
	var q quote
	var errChan = make(chan error)
	var wg sync.WaitGroup

	for key, url := range urls {
		wg.Add(1)
		go func(k, u string) {
			defer wg.Done()
			data, err := fetch.Fetch(u)
			if err != nil {
				errChan <- err
			}
			var r response
			switch k {
			case "price":
				if err := json.Unmarshal(data, &r); err != nil {
					errChan <- err
				}
				q.last, err = strconv.ParseFloat(r.Data.Amount, 64)
				if err != nil {
					errChan <- errors.Wrap(err, "failed to parse last to float")
				}
			case "bid":
				if err := json.Unmarshal(data, &r); err != nil {
					errChan <- err
				}
				q.bid, err = strconv.ParseFloat(r.Data.Amount, 64)
				if err != nil {
					errChan <- errors.Wrap(err, "failed to parse bid to float")
				}
			case "ask":
				if err := json.Unmarshal(data, &r); err != nil {
					errChan <- err
				}
				q.ask, err = strconv.ParseFloat(r.Data.Amount, 64)
				if err != nil {
					errChan <- errors.Wrap(err, "failed to parse ask to float")
				}
			}
		}(key, url)
	}
	wg.Wait()
	return &q, nil
}
