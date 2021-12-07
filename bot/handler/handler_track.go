package handler

import (
	"fmt"
	"time"

	"github.com/artnoi43/fngobot/bot"
)

// Track periodically calls SendQuote()
func (h *Handler) Track(s []bot.Security, r int, conf Config) {
	ticker := time.NewTicker(time.Duration(conf.TrackSeconds) * time.Second)
	c := 0

	// First quote right away
	h.SendQuote(s)

	for c < r-1 {
		select {
		case <-h.quit:
			h.notifyStop()
			return
		case <-ticker.C:
			h.SendQuote(s)
			c++
		}
	}
	h.send(fmt.Sprintf("Tracking done for %s", h.uuid))
}
