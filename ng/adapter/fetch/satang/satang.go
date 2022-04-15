package satang

import (
	"strconv"
)

type fetcher struct{}

func New() interface{} {
	return &fetcher{}
}

// Enum for parse() */
const (
	bid = iota
	ask
)

func parse(q *quote, val interface{}, bidAsk int) error {
	for k, v := range val.(map[string]interface{}) {
		switch k {
		case "price":
			price, err := strconv.ParseFloat(v.(string), 64)
			if err != nil {
				return err
			}
			switch bidAsk {
			case bid:
				q.bid = price
			case ask:
				q.ask = price
			}
		}
	}
	return nil
}
