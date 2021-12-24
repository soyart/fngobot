package handler

import (
	"errors"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/fetch"
)

// SendQuote sends quote(s) to users via chats.
// It is reused by tracking and alerting handlers.
func (h *handler) SendQuote(securities []bot.Security) {
	for _, s := range securities {
		q, err := s.Quote()
		if err != nil {
			var errMsg string
			if errors.Is(err, fetch.ErrNotFound) {
				errMsg = "Ticker %s not found"
			} else {
				errMsg = "Error getting %s quote"
			}
			errMsg = "ID: %s\n" + errMsg + " from %s"
			h.send(printer.Sprintf(errMsg,
				h.uuid, s.Tick, s.GetSrcStr()))
			return
		}
		str := "ID: %s\nQuote from %s\n%s\nBid: %f\nAsk: %f\nLast: %f\n"
		last, _ := q.Last()
		bid, _ := q.Bid()
		ask, _ := q.Ask()

		msg := printer.Sprintf(str,
			h.uuid, s.GetSrcStr(), s.Tick, bid, ask, last)
		h.send(msg)
	}
}
