package binance

import (
	"strconv"
	"strings"
	"sync"

	"github.com/artnoi43/fngobot/fetch"
	"github.com/pkg/errors"
)

const BASE_URL = "https://api.binance.com/api/v3"

type quote struct {
	last float64
	bid  float64
	ask  float64
}

func (q *quote) Last() (float64, error) {
	return q.last, nil
}
func (q *quote) Bid() (float64, error) {
	return q.bid, nil
}
func (q *quote) Ask() (float64, error) {
	return q.ask, nil
}

func Get(tick string) (fetch.Quoter, error) {
	info := info{
		symbol: tick + "USDT",
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

			// @TODO: figure out how to handle empty interface
			switch k {
			case "price":
				for key0, val0 := range data {
					switch key0 {
					case "price":
						last, err := strconv.ParseFloat(val0.(string), 64)
						if err != nil {
							errChan <- errors.Wrap(err, "failed to parse last to float")
						}
						q.last = last
					}
				}
			case "depth":
				for key0, val0 := range data {
					switch key0 {
					case "bidPrice":
						bid, err := strconv.ParseFloat(val0.(string), 64)
						if err != nil {
							errChan <- errors.Wrap(err, "failed to parse bid to float")
						}
						q.bid = bid
					case "askPrice":
						ask, err := strconv.ParseFloat(val0.(string), 64)
						if err != nil {
							errChan <- errors.Wrap(err, "failed to parse ask to float")
						}
						q.ask = ask
					}
				}
			}
		}(key, url)
	}
	go func() {
		wg.Wait()
		close(errChan)
	}()
	var errs []string
	for err := range errChan {
		errs = append(errs, err.Error())
	}
	if len(errs) > 0 {
		return nil, errors.Errorf("%v", strings.Join(errs, ", "))
	}
	return &q, nil
}

type info struct {
	symbol string
	url    string
}
