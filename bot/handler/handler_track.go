package handler

import (
	"fmt"
	"time"

	"github.com/artnoi43/fngobot/bot"
)

// Track periodically calls SendQuote()
func (h *handler) Track(s []bot.Security, r int, conf Config) {
	ticker := time.NewTicker(
		time.Duration(conf.TrackSeconds) * time.Second,
	)

	// First quote right away
	h.SendQuote(s)
	// r-1 bc 1st quote already sent
	for c := 0; c < r-1; {
		select {
		// Quit if received signal
		case <-h.Quit:
			h.notifyStop()
			return
		// Send quotes every N second
		case <-ticker.C:
			h.SendQuote(s)
			c++
		}
	}
	h.send(fmt.Sprintf("Tracking done for %s", h.Uuid))
}
