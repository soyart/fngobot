package clihandler

import (
	"time"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/bot/utils"
	"github.com/artnoi43/fngobot/config"
)

func (h *handler) PriceAlert(
	alert bot.Alert,
	conf *config.Config,
) {
	matchedChan := make(chan bool, conf.AlertTimes)
	errChan := make(chan error)
	ticker := time.NewTicker(
		time.Duration(conf.AlertInterval) * time.Second,
	)
	defer ticker.Stop()
	// First alert right away
	bot.GetQuoteAndAlert(&alert, matchedChan, errChan)
	// Then we range over the channels
	c := 0
	for c < conf.AlertTimes {
		select {
		case <-h.Quit:
			h.notifyStop()
			return
		case <-errChan:
			h.notifyStop()
			return
		case <-ticker.C:
			bot.GetQuoteAndAlert(&alert, matchedChan, errChan)
		case matched := <-matchedChan:
			if matched {
				utils.Printer.Printf(
					"[%s]\nALERT!\n%s (%s, %s) is now %s %f\n",
					h.UUID(),
					alert.Security.Tick,
					alert.GetSrcStr(),
					alert.GetQuoteTypeStr(),
					alert.GetCondStr(),
					alert.Target,
				)
				// Also send quote to user
				h.Quote([]bot.Security{
					alert.Security,
				})
				// Increment counter
				c++
			}
		}
	}
}