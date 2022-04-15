package usecase

import (
	"github.com/artnoi43/fngobot/ng/internal/enums"
)

// Security is a struct storing info about how to get the quotes.
type Security struct {
	Tick    string    `json:"symbol" yaml:"Symbol"`
	Src     enums.Src `json:"source" yaml:"Source"`
	Fetcher Fetcher
}
