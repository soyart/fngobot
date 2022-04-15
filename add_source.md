# Adding a quote source to [FnGoBot](README.md)

We implement quote fetching in `adapter/fetch`. For a source to be usable in FnGoBot, it's implementation need to satisfy the `usecase.Fetcher` and `usecase.Quoter` and interfaces.

## 1. `usecase.Fetcher`
`Fetcher` specifies 1 method: `Get(string) (Quoter, error)`, i.e. any structs that can return a `usecase.Quoter` (itself an alias to `entity.Quoter`).

## 2. `usecase.Quoter` or `entity.Quoter`
`Quoter` specifies 3 methods, all of which are of type `func(string) (float64, error)`. These methods are `Last`, `Bid`, and `Ask`, for getting last market price, best bid price, and best ask price respectively.

## 3. Package `adapter/fetch`
Create a new package in `fetch` package for the source.

## 4. Package `internal/enums`
After you are done with implementing `usecase` interfaces, update `source.go` with the new source enums.

## 5. Package `adapter/fetch`
After the new enum for the source is declared, go back to `adapter/fetch/fetch.go` and add your source fetcher and its `fetch.newFunc` function to map `newFetcherMap`.