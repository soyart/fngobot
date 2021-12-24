package handler

import (
	"encoding/json"
	"reflect"

	"github.com/artnoi43/fngobot/parse"
)

func (h *handler) SendHandlers() {
	var nullChecker = &parse.BotCommand{}
	if len(BotHandlers) > 1 {
		for _, h := range BotHandlers {
			if !reflect.DeepEqual(h.GetCmd(), nullChecker) {
				j, _ := json.MarshalIndent(h, "  ", "  ")
				h.send(string(j))
			}
		}
	} else {
		h.send("No handlers found")
	}
}