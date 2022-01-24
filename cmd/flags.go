package cmd

import (
	"flag"

	"github.com/artnoi43/fngobot/enums"
)

type Flags struct {
	ConfigFile string
}

func (f *Flags) Parse() {
	flag.StringVar(&f.ConfigFile, "c", enums.CONF, "Path to configuration file")
	flag.Parse()
}
