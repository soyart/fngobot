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

	// First alert right away
	bot.Match(&alert, matchedChan)

	c := 0
	for c < conf.AlertTimes {
		select {
		case <-h.Quit:
			h.notifyStop()
			return
		case <-ticker.C:
			bot.Match(&alert, matchedChan)
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
			} else {
				// @TODO: handle this case better (error getting quote)
				// (i.e. when false is sent to the channel)
				msg := printer.Sprintf(
					"[%s]\nError getting quote for %s from %s",
					h.UUID(),
					alert.Security.Tick,
					alert.GetSrcStr(),
				)
				h.send(msg)
			}
		}
	}
	// Alert user when done
	h.send(printer.Sprintf(
		"[%s]\nAlert done for %s",
		h.Uuid,
		alert.Security.Tick,
	))
	log.Printf("[%s]: Alert done\n", h.Uuid)
}
