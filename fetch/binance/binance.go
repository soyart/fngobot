package binance

import (
	"strconv"
	"strings"
	"sync"

	"github.com/artnoi43/fngobot/fetch"
	"github.com/pkg/errors"
)

const BaseURL = "https://api.binance.com"

type info struct {
	symbol string
	url    string
}

func Get(tick string) (fetch.Quoter, error) {
	urls := prepareURLs(info{
		symbol: tick + "USDT",
		url:    BaseURL,
	})
	var q quote
	var errChan = make(chan error)
	var wg sync.WaitGroup
	for key, url := range urls {
		wg.Add(1)
		go func(k, u string) {
			defer wg.Done()
			data, err := fetch.FetchMapStrInf(u)
			if err != nil {
				errChan <- err
			}

			// @TODO: refactor
			switch k {
			case "price":
				for key, val := range data {
					switch key {
					case "price":
						last, err := strconv.ParseFloat(val.(string), 64)
						if err != nil {
							errChan <- errors.Wrap(
								err,
								"failed to parse last to float",
							)
						}
						q.last = last
					}
				}
			case "depth":
				for key, val := range data {
					switch key {
					case "bidPrice":
						bid, err := strconv.ParseFloat(val.(string), 64)
						if err != nil {
							errChan <- errors.Wrap(
								err,
								"failed to parse bid to float",
							)
						}
						q.bid = bid
					case "askPrice":
						ask, err := strconv.ParseFloat(val.(string), 64)
						if err != nil {
							errChan <- errors.Wrap(
								err,
								"failed to parse ask to float",
							)
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
