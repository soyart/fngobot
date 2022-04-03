package tghandler

import (
	"fmt"
	"reflect"

	"github.com/go-yaml/yaml"

	"github.com/artnoi43/fngobot/lib/bot"
	"github.com/artnoi43/fngobot/lib/bot/handler/utils"
	"github.com/artnoi43/fngobot/lib/parse"
)

func (h *handler) SendHandlers() error {
	var nullChecker = &parse.BotCommand{}
	var runningHandlers []*handler
	for _, handlerInterface := range SenderHandlers[h.c.Message().Sender.ID] {
		// Discard null struct
		h, ok := handlerInterface.(*handler)
		if !ok {
			return fmt.Errorf("type assertion on Handler interface failed")
		}
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
		h.reply(msg)
		return nil
	}
	h.reply("No active handlers found")
	return nil
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
