package clihandler

import (
	"fmt"
	"log"
	"sync"

	"github.com/artnoi43/fngobot/lib/bot"
	"github.com/artnoi43/fngobot/lib/bot/handler/utils"
	"github.com/artnoi43/fngobot/lib/enums"
	"github.com/artnoi43/fngobot/lib/fetch"
)

func (h *handler) Quote(securities []bot.Security) {
	var wg sync.WaitGroup
	for _, security := range securities {
		wg.Add(1)
		go func(s bot.Security) {
			defer wg.Done()
			q, err := s.Quote()
			if err != nil {
				log.Printf(
					"Failed to fetch %s quote from %s: %s\n",
					s.Tick,
					s.Src,
					err.Error(),
				)
				return
			}
			printQuote(s.Tick, s.Src, q)
		}(security)
	}
	wg.Wait()
}

func printQuote(t string, s enums.Src, q fetch.Quoter) {
	bid, err := q.Bid()
	if err != nil {
		bid = -1
	}
	ask, err := q.Ask()
	if err != nil {
		ask = -1
	}
	last, err := q.Last()
	if err != nil {
		last = -1
	}
	utils.Printer.Printf(
		"Ticker: %s [%s]\nBid: %f\nAsk: %f\nLast: %f\n",
		t, s, bid, ask, last,
	)
	fmt.Println(enums.Bar)
}
