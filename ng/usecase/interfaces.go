package usecase

import "github.com/artnoi43/fngobot/ng/entity"

// For adapter/fetch to import
type Quoter entity.Quoter

// For adapter/fetch to import
type Fetcher interface {
	Get(tick string) (Quoter, error)
}
