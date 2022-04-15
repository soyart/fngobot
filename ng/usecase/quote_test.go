package usecase

import (
	"errors"
	"testing"

	"github.com/artnoi43/fngobot/ng/internal/enums"
)

func TestQuote(t *testing.T) {
	guard := make(chan struct{}, 4)
	failInvalidSource := func() {
		t.Fatal("invalid source returned from Quote()")
	}
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
		{Tick: "ada", Src: enums.Coinbase},
	}
	for _, s := range securities {
		guard <- struct{}{}
		go func(security *Security) {
			_, err := security.Quote()
			if err != nil {
				if errors.Is(enums.ErrInvalidSrc, err) {
					failInvalidSource()
				} else {
					t.Errorf(
						"error getting quote for %s from %s: %v\n",
						security.Tick,
						security.GetSrcStr(),
						err.Error(),
					)
				}
			}
			<-guard
		}(s)
	}
}
