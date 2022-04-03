package parse

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/enums"
)

func TestGetSrc(t *testing.T) {
	type getSrcOut struct {
		idx int
		src enums.Src
	}
	tests := []struct {
		In       string
		Expected getSrcOut
	}{

		{
			In: "YAHOO",
			Expected: getSrcOut{
				idx: 1,
				src: enums.Yahoo,
			},
		},
		{
			In: "CRYPTO",
			Expected: getSrcOut{
				idx: 2,
				src: enums.YahooCrypto,
			},
		},
		{
			In: "SATANG",
			Expected: getSrcOut{
				idx: 2,
				src: enums.Satang,
			},
		},
		{
			In: "BITKUB",
			Expected: getSrcOut{
				idx: 2,
				src: enums.Bitkub,
			},
		},
		{
			In: "BINANCE",
			Expected: getSrcOut{
				idx: 2,
				src: enums.Binance,
			},
		},
		{
			In: "COINBASE",
			Expected: getSrcOut{
				idx: 2,
				src: enums.Coinbase,
			},
		},
	}
	for _, test := range tests {
		idx, src := getSrc(test.In)
		if test.Expected.idx != idx {
			t.Errorf("invalid idx for %s\n", test.In)
		}
		if test.Expected.src != src {
			t.Errorf("invalid src for %s\n", test.In)
		}
	}
}

// Test parsing UserCommand into BotCommand
func TestParse(t *testing.T) {
	for _, test := range tests {
		out, err := test.In.Parse()
		if err != 0 {
			t.Errorf("error parsing UserCommand: %+v\n", test.In)
		}
		inJSON, _ := json.MarshalIndent(test.In, "  ", "  ")
		outJSON, _ := json.MarshalIndent(out, "  ", "  ")
		report := func() {
			fmt.Printf("In: %s\nOut: %s\n", inJSON, outJSON)
		}
		switch test.In.Type {
		case enums.QuoteBot:
			if !reflect.DeepEqual(out.Quote.Securities, test.Expected.Quote.Securities) {
				t.Errorf("[/quote] invalid quote securities")
				report()
			}
		case enums.TrackBot:
			if !reflect.DeepEqual(out.Quote.Securities, test.Expected.Quote.Securities) {
				t.Errorf("[/track] invalid quote securities")
				report()
			}
			if out.Track.TrackTimes != test.Expected.Track.TrackTimes {
				t.Errorf("[/track] invalid track time")
				report()
			}
		case enums.AlertBot:
			if !reflect.DeepEqual(out.Alert, test.Expected.Alert) {
				if !reflect.DeepEqual(out.Alert.Security, test.Expected.Alert.Security) {
					t.Errorf("[/alert] invalid alert security")
					report()
				}
				if out.Alert.Src != test.Expected.Alert.Src {
					t.Errorf("[/alert] invalid alert source")
					report()
				}
				if out.Alert.Condition != test.Expected.Alert.Condition {
					t.Errorf("[/alert] invalid alert condition")
					report()
				}
				if out.Alert.QuoteType != test.Expected.Alert.QuoteType {
					t.Errorf("[/alert] invalid alert quote type")
					report()
				}
				if out.Alert.Target != test.Expected.Alert.Target {
					t.Errorf("invalid alert target")
					report()
				}
			}
		}
	}
}

type parseTest struct {
	In       UserCommand
	Expected BotCommand
}

var (
	tests = []parseTest{
		{
			In: UserCommand{
				Type: enums.QuoteBot,
				Text: "/quote gc=f"},
			Expected: BotCommand{
				Quote: quoteCommand{
					Securities: []bot.Security{
						{
							Tick: "GC=F",
							Src:  enums.Yahoo,
						},
					},
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.QuoteBot,
				Text: "/quote satang btc",
			},
			Expected: BotCommand{
				Quote: quoteCommand{
					Securities: []bot.Security{
						{
							Tick: "BTC",
							Src:  enums.Satang,
						},
					},
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.QuoteBot,
				Text: "/quote bitkub btc",
			},
			Expected: BotCommand{
				Quote: quoteCommand{
					Securities: []bot.Security{
						{
							Tick: "BTC",
							Src:  enums.Bitkub,
						},
					},
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.QuoteBot,
				Text: "/quote binance btc"},
			Expected: BotCommand{
				Quote: quoteCommand{
					Securities: []bot.Security{
						{
							Tick: "BTC",
							Src:  enums.Binance,
						},
					},
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.QuoteBot,
				Text: "/quote coinbase ada"},
			Expected: BotCommand{
				Quote: quoteCommand{
					Securities: []bot.Security{
						{
							Tick: "ADA",
							Src:  enums.Coinbase,
						},
					},
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.TrackBot,
				Text: "/track gc=f 2",
			},
			Expected: BotCommand{
				Track: trackCommand{
					quoteCommand: quoteCommand{
						Securities: []bot.Security{
							{
								Tick: "GC=F",
								Src:  enums.Yahoo,
							},
						},
					},
					TrackTimes: 2,
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.TrackBot,
				Text: "/track satang btc 69",
			},
			Expected: BotCommand{
				Track: trackCommand{
					quoteCommand: quoteCommand{
						Securities: []bot.Security{
							{
								Tick: "BTC",
								Src:  enums.Satang,
							},
						},
					},
					TrackTimes: 69,
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.TrackBot,
				Text: "/track bitkub btc ada 69",
			},
			Expected: BotCommand{
				Track: trackCommand{
					quoteCommand: quoteCommand{
						Securities: []bot.Security{
							{Tick: "BTC", Src: enums.Bitkub},
							{Tick: "ADA", Src: enums.Bitkub},
						},
					},
					TrackTimes: 69,
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.TrackBot,
				Text: "/track binance btc ada 69",
			},
			Expected: BotCommand{
				Track: trackCommand{
					quoteCommand: quoteCommand{
						Securities: []bot.Security{
							{Tick: "BTC", Src: enums.Binance},
							{Tick: "ADA", Src: enums.Binance},
						},
					},
					TrackTimes: 69,
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.TrackBot,
				Text: "/track coinbase btc ada 69",
			},
			Expected: BotCommand{
				Track: trackCommand{
					quoteCommand: quoteCommand{
						Securities: []bot.Security{
							{Tick: "BTC", Src: enums.Coinbase},
							{Tick: "ADA", Src: enums.Coinbase},
						},
					},
					TrackTimes: 69,
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.AlertBot,
				Text: "/alert gc=f > 0",
			},
			Expected: BotCommand{
				Alert: bot.Alert{
					Security: bot.Security{
						Tick: "GC=F",
						Src:  enums.Yahoo,
					},
					Condition: enums.Gt,
					QuoteType: enums.Last,
					Target:    0,
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.AlertBot,
				Text: "/alert gc=f bid > 0",
			},
			Expected: BotCommand{
				Alert: bot.Alert{
					Security: bot.Security{
						Tick: "GC=F",
						Src:  enums.Yahoo,
					},
					Condition: enums.Gt,
					QuoteType: enums.Bid,
					Target:    0,
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.AlertBot,
				Text: "/alert satang btc bid > 112",
			},
			Expected: BotCommand{
				Alert: bot.Alert{
					Security: bot.Security{
						Tick: "BTC",
						Src:  enums.Satang,
					},
					Condition: enums.Gt,
					QuoteType: enums.Bid,
					Target:    112,
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.AlertBot,
				Text: "/alert bitkub btc < 112",
			},
			Expected: BotCommand{
				Alert: bot.Alert{
					Security: bot.Security{
						Tick: "BTC",
						Src:  enums.Bitkub,
					},
					Condition: enums.Lt,
					QuoteType: enums.Last,
					Target:    112,
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.AlertBot,
				Text: "/alert bitkub btc bid > 112",
			},
			Expected: BotCommand{
				Alert: bot.Alert{
					Security: bot.Security{
						Tick: "BTC",
						Src:  enums.Bitkub,
					},
					Condition: enums.Gt,
					QuoteType: enums.Bid,
					Target:    112,
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.AlertBot,
				Text: "/alert binance btc bid > 112",
			},
			Expected: BotCommand{
				Alert: bot.Alert{
					Security: bot.Security{
						Tick: "BTC",
						Src:  enums.Binance,
					},
					Condition: enums.Gt,
					QuoteType: enums.Bid,
					Target:    112,
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.AlertBot,
				Text: "/alert coinbase btc bid > 112",
			},
			Expected: BotCommand{
				Alert: bot.Alert{
					Security: bot.Security{
						Tick: "BTC",
						Src:  enums.Coinbase,
					},
					Condition: enums.Gt,
					QuoteType: enums.Bid,
					Target:    112,
				},
			},
		},
		{
			In: UserCommand{
				Type: enums.AlertBot,
				Text: "/alert coinbase btc < 112",
			},
			Expected: BotCommand{
				Alert: bot.Alert{
					Security: bot.Security{
						Tick: "BTC",
						Src:  enums.Coinbase,
					},
					Condition: enums.Lt,
					QuoteType: enums.Last,
					Target:    112,
				},
			},
		},
	}
)
