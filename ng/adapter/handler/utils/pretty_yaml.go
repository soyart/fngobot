package utils

import (
	"github.com/artnoi43/fngobot/ng/adapter/handler"
	"gopkg.in/yaml.v2"
)

func Yaml[T handler.Handler](h T) string {
	thisHandler := handler.PrettyHandler{
		Uuid:  h.UUID(),
		Quote: h.GetCmd().Quote.Securities,
		Track: h.GetCmd().Track.Securities,
		Alert: handler.PrettyAlert{
			Security:  h.GetCmd().Alert.Security,
			Condition: h.GetCmd().Alert.GetCondStr(),
			Target:    h.GetCmd().Alert.Target,
		},
		Start: h.StartTime().Format(TimeFormat),
	}
	y, _ := yaml.Marshal(&thisHandler)
	return string(y)
}
