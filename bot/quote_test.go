package bot

import (
	"testing"

	"github.com/artnoi43/fngobot/enums"
)

func TestQuote(t *testing.T) {
	var securities = []*Security{
		// keep them ticks lowercase or mixed-case
		{Tick: "btc", Src: enums.Satang},
		{Tick: "btc", Src: enums.Bitkub},
		{Tick: "btc", Src: enums.YahooCrypto},
		{Tick: "ada", Src: enums.Satang},
		{Tick: "ada", Src: enums.Bitkub},
		{Tick: "ada", Src: enums.YahooCrypto},
		{Tick: "bbl.bk", Src: enums.Yahoo},
		{Tick: "gc=f", Src: enums.Yahoo},
	}
	for _, s := range securities {
		_, err := s.Quote()
		if err != nil {
			t.Errorf("error getting quote for %s from %s: %v\n", s.Tick, s.GetSrcStr(), err.Error())
		}
	}
}
