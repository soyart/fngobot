package satang

import (
	"encoding/json"
	"strconv"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
	"github.com/artnoi43/fngobot/usecase"
	"github.com/pkg/errors"
)

// Get fetches data from Satang JSON API,
// and parses the fetched JSON into Quote struct
func (f *fetcher) Get(tick string) (usecase.Quoter, error) {
	/* Documentation for Satang:
	 * https://docs.satangcorp.com/apis/public/orders */
	u, err := prepareURL(&common.FetchInfo{
		URL:    BaseURL,
		Symbol: tick,
	})
	if err != nil {
		return nil, err
	}
	data, err := common.Fetch(u.String())
	if err != nil {
		return nil, err
	}

	var resp satangResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal satang json")
	}
	var q common.ApiQuote
	for key, openOrders := range resp {
		best := openOrders[0]
		populateBidAsk(key, &best, &q)
	}
	return &q, nil
}

func populateBidAsk(key string, best *satangQuote, q *common.ApiQuote) error {
	var f *float64
	switch key {
	case "bid":
		f = &q.Bid
	case "ask":
		f = &q.Ask
	}

	val, err := strconv.ParseFloat(best.Price, 64)
	if err != nil {
		errors.Wrap(err, "failed to parse ask string to float")
	}
	*f = val
	return nil
}
