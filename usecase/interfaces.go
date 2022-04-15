package usecase

import "github.com/artnoi43/fngobot/entity"

type Quoter entity.Quoter

// For adapter/fetch to import
type Fetcher interface {
	Get(tick string) (Quoter, error)
}
