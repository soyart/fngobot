package handler

import (
	"encoding/json"
	"reflect"

	"github.com/artnoi43/fngobot/parse"
)

func (h *handler) SendHandlers() {
	var nullChecker = &parse.BotCommand{}
	var okHandlers = []Handler{}
	for _, h := range BotHandlers {
		if !reflect.DeepEqual(h.GetCmd(), nullChecker) {
			okHandlers = append(okHandlers, h)
		}
	}
	if len(okHandlers) > 0 {
		for _, h := range okHandlers {
			j, _ := json.MarshalIndent(h, "  ", "  ")
			h.send(string(j))
		}
	} else {
		h.send("No active handlers found")
	}
}