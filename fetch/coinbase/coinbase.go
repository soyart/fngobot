package coinbase

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/artnoi43/fngobot/fetch"
	"github.com/pkg/errors"
)

const BASE_URL = "https://api.coinbase.com"

type info struct {
	symbol string
	url    string
}

func Get(tick string) (fetch.Quoter, error) {
	info := info{
		symbol: tick,
		url:    BASE_URL,
	}
	urlMap := prepareURLs(info)
	var q quote
	var errChan = make(chan error)
	var wg sync.WaitGroup

	for key, url := range urlMap {
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

type response struct {
	Data priceData `json:"data"`
}

type priceData struct {
	Amount string `json:"amount"`
}
