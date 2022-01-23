package tghandler

import (
	"reflect"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/bot/utils"
	"github.com/artnoi43/fngobot/parse"
	"github.com/go-yaml/yaml"
)

func (h *handler) SendHandlers() {
	var nullChecker = &parse.BotCommand{}
	var runningHandlers Handlers
	for _, h := range SenderHandlers[h.Msg.Sender.ID] {
		// Discard null struct
		if !reflect.DeepEqual(h.GetCmd(), nullChecker) {
			if h.isRunning() {
				runningHandlers = append(runningHandlers, h)
			}
		}
	}
	if len(runningHandlers) > 0 {
		// Use a for loop to append YAML string
		var msg string
		for _, runningHandler := range runningHandlers {
			msg = msg + runningHandler.yaml()
		}
		h.send(msg)
		return
	}
	h.send("No active handlers found")
}

func (h *handler) yaml() string {
	// These types are only for marshaling YAML
	type prettyAlert struct {
		Security  bot.Security `yaml:"Security,omitempty"`
		Condition string       `yaml:"Condition,omitempty"`
		Target    float64      `yaml:"Target,omitempty"`
	}
	type prettyHandler struct {
		Uuid  string         `yaml:"UUID,omitempty"`
		Start string         `yaml:"Start,omitempty"`
		Quote []bot.Security `yaml:"Quote,omitempty"`
		Track []bot.Security `yaml:"Track,omitempty"`
		Alert prettyAlert    `yaml:"Alert,omitempty"`
	}
	thisHandler := prettyHandler{
		Uuid:  h.Uuid,
		Quote: h.GetCmd().Quote.Securities,
		Track: h.GetCmd().Track.Securities,
		Alert: prettyAlert{
			Security:  h.GetCmd().Alert.Security,
			Condition: h.GetCmd().Alert.GetCondStr(),
			Target:    h.GetCmd().Alert.Target,
		},
		Start: h.Start.Format(utils.TimeFormat),
	}
	y, _ := yaml.Marshal(&thisHandler)
	return string(y)
}
