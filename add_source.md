# Adding a quote source to [FnGoBot](README.md)

## Package `enums`
Update `source.go` with the new source enums.

## Package `fetch`
Create a new package in `fetch` package for the source.

From there, compose a new struct, for example, struct `quote`. Make sure `quote` implement the `fetch.Quoter` interface.

Lastly, you'll need a `Get` method. The function signature for `Get` must match `fetch.FetchFunc` type alias. Usually, `Get` fetches JSON data from remote APIs and parses the data into something useful. Sometimes, as with `fetch/yahoo` and `fetch/crypto`, `Get` is just a wrapper for other library functions.

## Package `bot`
In `quote.go`, populate the `quoteFuncs`. It's just a map of your sources (enums) and their `Get` methods.
