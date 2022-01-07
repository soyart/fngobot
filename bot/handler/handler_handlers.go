package handler

import (
	"reflect"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/parse"
	"github.com/go-yaml/yaml"
)

func (h *handler) SendHandlers() {
	var nullChecker = &parse.BotCommand{}
	var runningHandlers = []Handler{}
	for _, h := range BotHandlers {
		if !reflect.DeepEqual(h.GetCmd(), nullChecker) {
			// Add running handler to okHandlers
			if h.isRunning() {
				runningHandlers = append(runningHandlers, h)
			}
		}
	}
	if len(runningHandlers) > 0 {
		var s string
		for _, runningHandler := range runningHandlers {
			s = s + runningHandler.yaml()
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
		Start: h.Start.Format(timeFormat),
	}
	y, _ := yaml.Marshal(&thisHandler)
	return string(y)
}
