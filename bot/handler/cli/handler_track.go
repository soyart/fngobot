package clihandler

import (
	"strings"
	"sync"
	"time"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/bot/utils"
)

func (h *handler) Track(
	securities []bot.Security,
	r int, // Track rounds
) {
	utils.Printer.Printf(
		"Starting tracker: %dx, every %d sec\n\n",
		r, h.conf.Handler.TrackInterval,
	)
	// First quote fires right away (hence r-1)
	h.Quote(securities)
	quoteAll := func() {
		for c := 0; c < r-1; c++ {
			h.Quote(securities)
		}
	}

	ticker := time.NewTicker(
		time.Duration(h.conf.Handler.TrackInterval) * time.Second,
	)
	defer ticker.Stop()

	// This Goroutine fires quote every ticking second
	// and also listening on h.Quit
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ticker.C:
			quoteAll()
		// Impossible for now
		case <-h.Quit:
			h.notifyStop()
			return
		}
	}()
	wg.Wait()
	tickers := make([]string, len(securities))
	for i, security := range securities {
		tickers[i] = security.Tick
	}
	utils.Printer.Printf(
		"[%s] Tracking done for %s\n",
		h.UUID(),
		strings.Join(tickers, ", "),
	)
}
