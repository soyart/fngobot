package tghandler

import (
	"log"
	"time"

	"github.com/artnoi43/fngobot/lib/bot"
	"github.com/artnoi43/fngobot/lib/bot/handler/utils"
)

// PriceAlert sends alerts to users if a condition is matched.
func (h *handler) PriceAlert(alert bot.Alert) {
	// Notify user of the alert handler
	chatMsg := h.c.Message()
	startMsg := utils.Printer.Sprintf(
		"Your alert handler ID is %s\nMessage: %s\nTime: %s)",
		h.Uuid,
		chatMsg.Text,
		chatMsg.Time().Format(utils.TimeFormat),
	)
	h.reply(startMsg)
	// Channels for alerting and time ticker
	matchedChan := make(chan bool, h.conf.Handler.AlertTimes)
	errChan := make(chan error)
	ticker := time.NewTicker(
		time.Duration(h.conf.Handler.AlertInterval) * time.Second,
	)
	defer ticker.Stop()
	// First alert right away
	bot.GetQuoteAndAlert(&alert, matchedChan, errChan)
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
			bot.GetQuoteAndAlert(&alert, matchedChan, errChan)
		case matched := <-matchedChan:
			if matched {
				msg := utils.Printer.Sprintf(
					"[%s]\nALERT!\n%s (%s, %s) is now %s %f",
					h.UUID(),
					alert.Security.Tick,
					alert.GetSrcStr(),
					alert.GetQuoteTypeStr(),
					alert.GetCondStr(),
					alert.Target,
				)
				h.reply(msg)
				// Also reply quote to user
				h.Quote([]bot.Security{
					alert.Security,
				})
				// Increment counter
				c++
			}
		}
	}
	// Alert user when done
	h.reply(utils.Printer.Sprintf(
		"[%s]: Alert Finished %s (%s)\n",
		h.Uuid,
		alert.Security.Tick,
		alert.GetSrcStr(),
	))
	log.Printf(
		"[%s]: Alert Finished %s (%s)\n",
		h.Uuid,
		alert.Security.Tick,
		alert.GetSrcStr(),
	)
}
