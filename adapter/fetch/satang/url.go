package satang

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
)

// const BaseURL = "https://satangcorp.com/api/orderbook-tickers/"
const BaseURL = "https://satangcorp.com"

func prepareURL(info *common.FetchInfo) (*url.URL, error) {
	u, err := url.Parse(info.URL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse %s into url.URL", BaseURL)
	}
	u.Path = "/api/orders/"
	u.RawQuery = fmt.Sprintf("pair=%s_THB", info.Symbol)
	return u, nil
}
