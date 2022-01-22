## github.com/artnoi43/fngobot/fetch/binance
File `binance.go` defines `Get()` function, while `quoter.go` defines a struct `quote` implementing `fetch.Quoter` interface.

`Get()` fetches the API data from Coinbase API in JSON, parses that JSON data into a Go object (struct `quote`) before returning the address of that object *if* the given ticker symbol `tick` is valid.

If the ticker symbol is invalid and cannot be found in the JSON data, or an error was encountered, `Get()` returns `nil` and a custom error.