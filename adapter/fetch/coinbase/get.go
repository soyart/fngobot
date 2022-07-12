package coinbase

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/pkg/errors"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
	"github.com/artnoi43/fngobot/usecase"
)

func (f *fetcher) Get(tick string) (usecase.Quoter, error) {
	urls := prepareURLs(&common.FetchInfo{
		URL:    BaseURL,
		Symbol: tick,
	})
	var q common.ApiQuote
	var errChan = make(chan error)
	var wg sync.WaitGroup

	for key, url := range urls {
		wg.Add(1)
		go func(k, u string) {
			defer wg.Done()
			data, err := common.Fetch(u)
			if err != nil {
				errChan <- err
			}
			var r response
			switch k {
			case "price":
				if err := json.Unmarshal(data, &r); err != nil {
					errChan <- err
				}
				q.Last, err = strconv.ParseFloat(r.Data.Amount, 64)
				if err != nil {
					errChan <- errors.Wrap(err, "failed to parse last to float")
				}
			case "bid":
				if err := json.Unmarshal(data, &r); err != nil {
					errChan <- err
				}
				q.Bid, err = strconv.ParseFloat(r.Data.Amount, 64)
				if err != nil {
					errChan <- errors.Wrap(err, "failed to parse bid to float")
				}
			case "ask":
				if err := json.Unmarshal(data, &r); err != nil {
					errChan <- err
				}
				q.Ask, err = strconv.ParseFloat(r.Data.Amount, 64)
				if err != nil {
					errChan <- errors.Wrap(err, "failed to parse ask to float")
				}
			}
		}(key, url)
	}
	wg.Wait()
	return &q, nil
}
