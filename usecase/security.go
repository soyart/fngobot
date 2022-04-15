package usecase

import (
	"github.com/artnoi43/fngobot/internal/enums"
)

// Security is a struct storing info about how to get the quotes.
// Everything is exported for package adapter to construct this struct.
type Security struct {
	Tick    string    `json:"symbol" yaml:"Symbol"`
	Src     enums.Src `json:"source" yaml:"Source"`
	Fetcher Fetcher
}
