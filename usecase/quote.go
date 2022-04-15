package usecase

import (
	"strings"

	"github.com/artnoi43/fngobot/entity"
	"github.com/artnoi43/fngobot/internal/enums"
)

// Quote quotes Security instance. Quote sources are identified by s.Src.
// This function is called by many almost all handlers.
func (s *Security) Quote() (q entity.Quoter, err error) {
	if s.Src.IsValid() {
		s.Tick = strings.ToUpper(s.Tick)
		q, err = s.Fetcher.Get(s.Tick)
		if err != nil {
			return nil, err
		}
		return q, nil
	}
	// Should not happen
	// since parsing defaults to Yahoo Finance
	return nil, enums.ErrInvalidSrc
}
