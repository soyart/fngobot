package bot

import (
	"github.com/artnoi43/fngobot/enums"
	"github.com/go-yaml/yaml"
)

// Security is a struct storing info about how to get the quotes.
type Security struct {
	Tick string    `yaml:"Ticker"`
	Src  enums.Src `yaml:"Source"`
}

func (s Security) Yaml() string {
	y, _ := yaml.Marshal(&s)
	return string(y)
}
