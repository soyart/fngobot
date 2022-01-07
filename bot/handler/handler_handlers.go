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
		Uuid  string `yaml:"UUID,omitempty"`
		Start string `yaml:"Start,omitempty"`
		//		Q     []bot.Security `yaml:"Quote,omitempty"`
		//		T     []bot.Security `yaml:"Track,omitempty"`
		//		A     bot.Security   `yaml:"Alert,omitempty"`
		Quote []bot.Security `yaml:"Quote,omitempty"`
		Track []bot.Security `yaml:"Track,omitempty"`
		Alert bot.Security   `yaml:"Alert,omitempty"`
		Src   string         `yaml:"Source,omitempty"`
	}

	//	var quoteStrings []string
	//	var trackStrings []string
	//	var alertString string
	var src string

	if len(h.Cmd.Quote.Securities) > 0 {
		//		for _, security := range h.Cmd.Quote.Securities {
		//			quoteStrings = append(quoteStrings, security.Yaml())
		//		}
		src = string(h.Cmd.Quote.Securities[0].Src)
	}
	if len(h.Cmd.Track.Securities) > 0 {
		//		for _, security := range h.Cmd.Track.Securities {
		//			trackStrings = append(trackStrings, security.Yaml())
		//		}
		src = string(h.Cmd.Track.Securities[0].Src)
	}
	if len(h.Cmd.Alert.Tick) > 0 {
		//		alertString = h.Cmd.Alert.Yaml()
		src = string(h.Cmd.Alert.Src)
	}
	prettyYaml := prettyHandler{
		Uuid:  h.Uuid,
		Quote: h.Cmd.Quote.Securities,
		Track: h.Cmd.Track.Securities,
		Alert: h.Cmd.Alert.Security,
		Src:   src,
		Start: h.Start.Format("2006-01-02 15:04:05"),
	}
	y, _ := yaml.Marshal(&prettyYaml)
	return string(y)
}
