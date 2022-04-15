package tghandler

import (
	"fmt"
	"strings"
	"time"

	"github.com/artnoi43/fngobot/ng/usecase"
)

// Track periodically calls SendQuote()
func (h *handler) Track(
	securities []usecase.Security,
	r int, // Track times
) {
	ticker := time.NewTicker(
		time.Duration(h.conf.Handler.TrackInterval) * time.Second,
	)
	defer ticker.Stop()

	// First quote right away
	h.Quote(securities)
	// r-1 bc 1st quote already sent
	for c := 0; c < r-1; {
		select {
		// Quit if received signal
		case <-h.Quit:
			h.notifyStop()
			return
		// Send quotes every N second
		case <-ticker.C:
			h.Quote(securities)
			c++
		}
	}
	tickers := make([]string, len(securities))
	for i, security := range securities {
		tickers[i] = security.Tick
	}
	h.reply(fmt.Sprintf("[%s]\nTracking done for %s", h.Uuid, strings.Join(tickers, ", ")))
}
