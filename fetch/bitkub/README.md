## github.com/artnoi43/fngobot/fetch/bitkub
The file `bitkub.go` defines `Get()` function and struct `Quote`.

`Quote` currently has 6 fields: `Last` (JSON: last), `Bid` (JSON: highestBid), `Ask` (JSON: lowestAsk), `High` (JSON: high24hr), `Low` (JSON: low24hr), and `Change` (JSON: percentageChange).

`Get()` fetches the API data from Bitkub.com in JSON, parses that JSON data into a Go object (struct `Quote`) before returning the pointer to that object *if* the given ticker symbol `tick` is valid.

If the ticker symbol is invalid and cannot be found in the JSON data, or an error was encountered, `Get()` returns `nil` and a custom error.

## Example

    package main

	import (
		"fmt"
		"github.com/artnoi43/fngobot/api"
		"github.com/artnoi43/fngobot/api/bitkub"
	)

	func main() {
		tick := "BTC"
		q, err := bitkub.Get(tick)
		if err == api.NotFound {
			fmt.Printf("Ticker not found in JSON", tick)
			panic(err)
		} else if err != nil {
			panic(err)
		}

		fmt.Printf("Current quote on %s:\nLast: %f\nBid: %f\nAsk: %f\n",
			tick, q.Last, q.Bid, q.Ask)
	}

## JSON API

    {
		"error":0,
		"data":{
			"THB_AAVE":{
				"id":75,
				"name":"Aave",
				"alias":"AAVE",
				"last":8378,
				"lowestAsk":8352.1,
				"highestBid":8301.54,
				"percentChange":15.22,
				"baseVolume":1089.57180734,
				"quoteVolume":8515672.23,
				"isFrozen":0,
				"high24hr":8450,
				"low24hr":7222.61,
				"graph":[7271.37,7408,7352,7359.62,7311.94,7231.14,7284.03,7345.56,7436.7,7468.92,7422.94,7435.91,7484.53,7388.44,7406.06,7506,7470,7944.16,8079,8288,8371,8300,8243,8378],
				"order":75
			},
			"THB_ABT":{
				"id":22,
				"name":"Arcblock",
				"alias":"ABT",
				"last":4.12,
				"lowestAsk":4.12,
				"highestBid":4.07,
				"percentChange":13.5,
				"baseVolume":3365813.68514452,
				"quoteVolume":13268332.67,
				"isFrozen":0,
				"high24hr":4.29,
				"low24hr":3.44,
				"graph":[3.64,3.62,3.53,3.72,4.09,4.06,3.89,4.26,4.09,3.98,4.02,4.09,3.9,3.95,3.96,4.1,4.05,4.12,4.07,4.09,4.06,4.01,4.02,4.03,4.12],
				"order":24
			}
		}
	}
