package bot

import (
	"sync"
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
		{Tick: "btc", Src: enums.Binance},
	}
	var wg sync.WaitGroup
	for _, s := range securities {
		wg.Add(1)
		go func(security *Security) {
			defer wg.Done()
			_, err := security.Quote()
			if err != nil {
				t.Errorf(
					"error getting quote for %s from %s: %v\n",
					security.Tick,
					security.GetSrcStr(),
					err.Error(),
				)
			}
		}(s)
	}
	wg.Wait()
}
