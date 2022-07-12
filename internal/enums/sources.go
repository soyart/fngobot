package enums

import "fmt"

type (
	Src    int
	Switch string
)

var ErrInvalidSrc = fmt.Errorf("invalid source")

// quote sources - when adding new sources,
// also add them to validSrc below
const (
	Yahoo Src = iota
	YahooCrypto
	Satang
	Bitkub
	Binance
	Coinbase

	YahooSw       Switch = ""
	YahooCryptoSw Switch = "CRYPTO"
	SatangSw      Switch = "SATANG"
	BitkubSw      Switch = "BITKUB"
	BinanceSw     Switch = "BINANCE"
	CoinbaseSw    Switch = "COINBASE"
)

var validSrc = []Src{
	Yahoo,
	YahooCrypto,
	Satang,
	Bitkub,
	Binance,
	Coinbase,
}

// SwitchSrcMap -> The int part is the index of the word that represent
// start of the security names. The first index (0) is always the command string.
// For example, consider the following commands:
// /quote bbl.bk             => index = 1
// /quote coinbase bbl.bk    => index = 2
var SwitchSrcMap = map[Switch]map[Src]int{
	YahooSw:       {Yahoo: 1},
	YahooCryptoSw: {YahooCrypto: 2},
	SatangSw:      {Satang: 2},
	BitkubSw:      {Bitkub: 2},
	BinanceSw:     {Binance: 2},
	CoinbaseSw:    {Coinbase: 2},
}

func (s Src) String() string {
	switch s {
	case Yahoo:
		return "Yahoo! Finance"
	case YahooCrypto:
		return "Yahoo! Finance Cryptocurrency"
	case Satang:
		return "Satang Pro"
	case Bitkub:
		return "Bitkub"
	case Binance:
		return "Binance"
	case Coinbase:
		return "Coinbase"
	}
	return "Invalid quote source"
}
