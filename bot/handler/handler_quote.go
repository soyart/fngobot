package handler

import (
	"errors"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/fetch"
)

// SendQuote sends quote(s) to users via chats.
// It is reused by tracking and alerting handlers.
func (h *handler) SendQuote(securities []bot.Security) {
	for _, security := range securities {
		q, err := security.Quote()
		if err != nil {
			var errMsg string
			if errors.Is(err, fetch.ErrNotFound) {
				errMsg = "Ticker not found"
			} else {
				errMsg = "Error getting quote"
			}
			h.send(printer.Sprintf(
				"[%s]\n%s: %s from %s",
				h.UUID(),
				errMsg,
				security.Tick,
				security.GetSrcStr(),
			))
			return
		}
		last, _ := q.Last()
		bid, _ := q.Bid()
		ask, _ := q.Ask()
		msg := printer.Sprintf(
			"[%s]\nQuote from %s\n%s\nBid: %f\nAsk: %f\nLast: %f\n",
			h.UUID(),
			security.GetSrcStr(),
			security.Tick,
			bid,
			ask,
			last,
		)
		h.send(msg)
	}
}
