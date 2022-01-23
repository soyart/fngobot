package clihandler

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/config"
)

func (h *handler) Track(
	securities []bot.Security,
	r int, // Track rounds
	conf *config.Config,
) {
	var wg sync.WaitGroup
	ticker := time.NewTicker(
		time.Duration(conf.TrackSeconds) * time.Second,
	)
	for _, security := range securities {
		wg.Add(1)
		go func(s bot.Security) {
			// First quote right away
			h.Quote([]bot.Security{s})
			// r-1 bc 1st quote already sent
			for c := 0; c < r-1; {
				select {
				// Quit if received signal
				case <-h.Quit:
					h.notifyStop()
					return
				// Send quotes every N second
				case <-ticker.C:
					h.Quote([]bot.Security{s})
					c++
				}
			}
			fmt.Printf(
				"Tracking done for %s\n",
				h.Uuid,
			)
			os.Exit(0)
		}(security)
	}
	wg.Wait()
}
