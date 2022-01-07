package handler

import (
	"reflect"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/parse"
	"github.com/go-yaml/yaml"
)

func (h *handler) SendHandlers() {
	var nullChecker = &parse.BotCommand{}
	var okHandlers = []Handler{}
	for _, h := range BotHandlers {
		if !reflect.DeepEqual(h.GetCmd(), nullChecker) {
			// Add running handler to okHandlers
			if h.isRunning() {
				okHandlers = append(okHandlers, h)
			}
		}
	}
	if len(okHandlers) > 0 {
		var s string
		for _, okHandler := range okHandlers {
			s = s + okHandler.yaml()
		}
		h.send(s)
		return
	}
	h.send("No active handlers found")
}

func (h *handler) yaml() string {
	// This type is only for marshaling YAML
	type prettyHandler struct {
		Uuid  string         `yaml:"UUID,omitempty"`
		Start string         `yaml:"Start,omitempty"`
		Quote []bot.Security `yaml:"Quote,omitempty"`
		Track []bot.Security `yaml:"Track,omitempty"`
		Alert bot.Security   `yaml:"Alert,omitempty"`
	}
	thisHandler := prettyHandler{
		Uuid:  h.Uuid,
		Quote: h.Cmd.Quote.Securities,
		Track: h.Cmd.Track.Securities,
		Alert: h.Cmd.Alert.Security,
		Start: h.Start.Format("2006-01-02 15:04:05"),
	}
	y, _ := yaml.Marshal(&thisHandler)
	return string(y)
}
