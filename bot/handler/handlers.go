package handler

import "log"

type Handlers []*Handler

func (h *Handlers) Stop(uuid string) {
	for _, handler := range *h {
		switch uuid {
		case handler.uuid:
			log.Printf("[%s]: Sending quit signal\n", handler.uuid)
			handler.quit <- true
			log.Printf("[%s]: Sent quit signal\n", handler.uuid)
		}
	}
}