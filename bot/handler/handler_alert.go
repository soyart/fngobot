package handler

import (
	"log"
	"time"

	"github.com/artnoi43/fngobot/bot"
)

// PriceAlert sends alerts to users if a condition is matched.
func (h *Handler) PriceAlert(a bot.Alert, conf Config) {
	// Notify user of the handler
	startMsg := printer.Sprintf("Your alert handler ID is %s\nFrom message: %s at %v",
		h.uuid, h.msg.Text, h.msg.Time())
	h.send(startMsg)

	ticker := time.NewTicker(time.Duration(conf.AlertInterval) * time.Second)
	matched := make(chan bool, conf.AlertTimes)

	// First alert right away
	a.Match(matched)

	c := 0
	for c < conf.AlertTimes {
		select {
		case <-h.quit:
			h.notifyStop()
			return
		case <-ticker.C:
			a.Match(matched)
		case m := <-matched:
			if m {
				str := "ID: %s\nALERT!\n%s (%s) is now %s %f\non %s"
				msg := printer.Sprintf(str,
					h.uuid, a.Security.Tick, a.GetQuoteTypeStr(), a.GetCondStr(), a.Target, a.GetSrcStr())
				h.send(msg)
				// Also send quote to user
				h.SendQuote([]bot.Security{a.Security})
				c++
			}
		}
	}
	// Alert user when done
	doneStr := "ID: %s\nAlert done for %s"
	h.send(printer.Sprintf(doneStr, h.uuid, a.Security.Tick))
	log.Printf("[%s]: Alert done\n", h.uuid)
}
