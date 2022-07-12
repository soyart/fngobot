package tghandler

import (
	"errors"
	"sync"

	"github.com/artnoi43/fngobot/adapter/fetch/common"
	"github.com/artnoi43/fngobot/adapter/handler/utils"
	"github.com/artnoi43/fngobot/usecase"
)

// Quote is called by other handler methods
// to display security quotes to users.
func (h *handler) Quote(securities []usecase.Security) {
	// quotes := make(chan fetch.Quoter, len(securities))
	var wg sync.WaitGroup
	for _, security := range securities {
		wg.Add(1)
		// This Goroutines get quotes
		go func(s usecase.Security) {
			defer wg.Done()
			q, err := s.Quote()
			if err != nil {
				var errMsg string
				if errors.Is(err, common.ErrNotFound) {
					errMsg = "ticker not found"
				} else {
					errMsg = err.Error()
				}
				h.reply(utils.Printer.Sprintf(
					"[%s]\n%s: %s from %s",
					h.UUID(),
					errMsg,
					s.Tick,
					s.GetSrcStr(),
				))
				return
			}
			last, _ := q.QuoteLast()
			bid, _ := q.QuoteBid()
			ask, _ := q.QuoteAsk()
			msg := utils.Printer.Sprintf(
				"[%s]\nQuote from %s\n%s\nBid: %f\nAsk: %f\nLast: %f\n",
				h.UUID(), s.GetSrcStr(), s.Tick, bid, ask, last,
			)
			h.reply(msg)
		}(security)
	}
	wg.Wait()
}
