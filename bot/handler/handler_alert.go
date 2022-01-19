package handler

import (
	"log"
	"time"

	"github.com/artnoi43/fngobot/bot"
)

// PriceAlert sends alerts to users if a condition is matched.
func (h *handler) PriceAlert(alert bot.Alert, conf Config) {
	// Notify user of the handler
	startMsg := printer.Sprintf(
		"Your alert handler ID is %s\nMessage: %s\nTime: %s)",
		h.Uuid,
		h.Msg.Text,
		h.Msg.Time().Format(timeFormat),
	)
	h.send(startMsg)

	ticker := time.NewTicker(time.Duration(conf.AlertInterval) * time.Second)
	matchedChan := make(chan bool, conf.AlertTimes)
	errChan := make(chan error)
	// First alert right away
	bot.GetQuoteAndAlert(&alert, matchedChan, errChan)

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
		case m := <-matchedChan:
			if m {
				msg := printer.Sprintf(
					"[%s]\nALERT!\n%s (%s) is now %s %f\non %s",
					h.UUID(),
					alert.Security.Tick,
					alert.GetQuoteTypeStr(),
					alert.GetCondStr(),
					alert.Target,
					alert.GetSrcStr(),
				)
				h.send(msg)
				// Also send quote to user
				h.SendQuote([]bot.Security{
					alert.Security,
				})
				// Increment counter
				c++
			}
		}
	}
	// Alert user when done
	h.send(printer.Sprintf(
		"[%s]\nAlert done for %s",
		h.Uuid,
		alert.Security.Tick,
	))
	log.Printf(
		"[%s]: Alert done for %s (%s)\n",
		h.Uuid,
		alert.Security.Tick,
		alert.GetSrcStr(),
	)
}
