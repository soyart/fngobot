## github.com/artnoi43/fngobot/fetch/satang
The file `satang.go` defines `Get()` function struct `Quote`

`Quote` currently has 2 `float64` fields: `Bid` and `Ask`, which are best bid and ask prices respectively.

`Get()` fetches the API data from Satangcorp.com in JSON, parses that JSON data into a Go object (struct `Quote`) before returning the pointer to that object *if* the given ticker symbol `tick` is valid.

If the ticker symbol is invalid and cannot be found in the JSON data, or an error was encountered, `Get()` returns `nil` and a custom error.

## Example

    package main

	import (
		"fmt"
		"github.com/artnoi43/fngobot/api"
		"github.com/artnoi43/fngobot/api/satang"
	)

	func main() {
		tick := "BTC"
		q, err := satang.Get(tick)
		if err == api.NotFound {
			fmt.Printf("%s not found in JSON", tick)
			panic(err)
		} else if err ! nil {
			panic(err)
		}

		fmt.Printf("Current quote on %s:\nBid: %f\nAsk: %f\n",
			tick, q.Bid, q.Ask)
	}

## JSON API

    {
		"ADA_THB":{
			"bid":{
				"price":"44.66",
				"amount":"44.6"
			},
			"ask":{
				"price":"45",
				"amount":"107.2"
			}
		},
		"ALGO_THB":{
			"bid":{
				"price":"27.5",
				"amount":"122.64"
			},
			"ask":{
				"price":"28.5",
				"amount":"1200"
			}
		},
		"ATOM_THB":{
			"bid":{
				"price":"376",
				"amount":"0.64"
			},
			"ask":{
				"price":"383.99",
				"amount":"2.93"
			}
		},
		"BAND_THB":{
			"bid":{
				"price":"170.01",
				"amount":"3.39"
			},
			"ask":{
				"price":"185",
				"amount":"9.96"
			}
		}
	}
