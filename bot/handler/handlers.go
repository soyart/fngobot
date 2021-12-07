package handler

import "log"

// Handlers is a collection of handlers, used for stopping a handler
type Handlers []*Handler

// Stop stops a handler with matching UUID
func (h *Handlers) Stop(uuid string) (i int, ok bool) {
	for idx, handler := range *h {
		switch uuid {
		case handler.uuid:
			log.Printf("[%s]: Sending quit signal\n", handler.uuid)
			handler.quit <- true
			log.Printf("[%s]: Sent quit signal\n", handler.uuid)
			i = idx
			ok = true
		}
	}
	return i, ok
}
