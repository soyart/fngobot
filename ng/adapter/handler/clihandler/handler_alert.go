package clihandler

import (
	"time"

	"github.com/artnoi43/fngobot/ng/adapter/handler/utils"
	"github.com/artnoi43/fngobot/ng/usecase"
)

func (h *handler) PriceAlert(
	alert usecase.Alert,
) {
	matchedChan := make(chan bool, h.conf.Handler.AlertTimes)
	errChan := make(chan error)
	ticker := time.NewTicker(
		time.Duration(h.conf.Handler.AlertInterval) * time.Second,
	)
	defer ticker.Stop()
	// First alert right away
	usecase.GetQuoteAndAlert(&alert, matchedChan, errChan)
	// Then we range over the channels
	c := 0
	for c < h.conf.Handler.AlertTimes {
		select {
		case <-h.Quit:
			h.notifyStop()
			return
		case <-errChan:
			h.notifyStop()
			return
		case <-ticker.C:
			usecase.GetQuoteAndAlert(&alert, matchedChan, errChan)
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
				h.Quote([]usecase.Security{
					alert.Security,
				})
				// Increment counter
				c++
			}
		}
	}
	utils.Printer.Printf(
		"[%s] Alert Finished", h.UUID(),
	)
	h.done <- struct{}{}
}
