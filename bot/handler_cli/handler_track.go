package clihandler

import (
	"os"
	"sync"
	"time"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/bot/utils"
	"github.com/artnoi43/fngobot/config"
)

func (h *handler) Track(
	securities []bot.Security,
	r int, // Track rounds
	conf *config.Config,
) {
	utils.Printer.Printf(
		"Starting tracker: %dx, every %d sec\n\n",
		r, conf.TrackInterval,
	)
	// First quote fires right away (hence r-1)
	h.Quote(securities)
	quoteAll := func() {
		for c := 0; c < r-1; c++ {
			h.Quote(securities)
		}
	}
	ticker := time.NewTicker(
		time.Duration(conf.TrackInterval) * time.Second,
	)

	// This Goroutine fires quote every ticking second
	// and also listening on h.Quit
	defer ticker.Stop()
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
			os.Exit(0)
		}
	}()
	wg.Wait()
	utils.Printer.Printf(
		"[%s] Tracking done\n",
		h.UUID(),
	)
}
