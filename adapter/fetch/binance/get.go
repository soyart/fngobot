package binance

import (
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
	"github.com/artnoi43/fngobot/internal/enums"
	"github.com/artnoi43/fngobot/usecase"
)

func (f *fetcher) Get(tick string) (usecase.Quoter, error) {
	urls := prepareURLs(&common.FetchInfo{
		URL:    BaseURL,
		Symbol: tick + "USDT",
	})
	var q common.ApiQuote
	var errChan = make(chan error)
	var wg sync.WaitGroup
	for key, url := range urls {
		wg.Add(1)
		go func(k, u string) {
			defer wg.Done()
			data, err := common.FetchMapStrInf(u)
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
						q.Last = last
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
						q.Bid = bid
					case "askPrice":
						ask, err := strconv.ParseFloat(val.(string), 64)
						if err != nil {
							errChan <- errors.Wrap(
								err,
								"failed to parse ask to float",
							)
						}
						q.Ask = ask
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

	q.Src = enums.Binance
	return &q, nil
}
