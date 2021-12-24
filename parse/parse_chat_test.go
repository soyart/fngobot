package parse

import (
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
		In  string
		Out getSrcOut
	}{
		{
			In: "CRYPTO",
			Out: getSrcOut{
				idx: 2,
				src: enums.YahooCrypto,
			},
		},
		{
			In: "SATANG",
			Out: getSrcOut{
				idx: 2,
				src: enums.Satang,
			},
		},
		{
			In: "BITKUB",
			Out: getSrcOut{
				idx: 2,
				src: enums.Bitkub,
			},
		},
		{
			In: "YAHOO",
			Out: getSrcOut{
				idx: 1,
				src: enums.Yahoo,
			},
		},
	}
	for _, test := range tests {
		idx, src := getSrc(test.In)
		if test.Out.idx != idx {
			t.Errorf("invalid idx for %s\n", test.In)
		}
		if test.Out.src != src {
			t.Errorf("invalid src for %s\n", test.In)
		}
	}
}

// Test parsing UserCommand into BotCommand
func TestParse(t *testing.T) {
	type parseTest struct {
		In  UserCommand
		Out BotCommand
	}

	tests := []parseTest{
		{
			In: UserCommand{
				Command: QuoteCmd,
				Chat:    "/quote gc=f"},
			Out: BotCommand{
				Quote: quoteCommand{
					Securities: []bot.Security{
						{Tick: "gc=f", Src: enums.Yahoo},
					},
				},
			},
		},
		{
			In: UserCommand{
				Command: QuoteCmd,
				Chat:    "/quote satang btc",
			},
			Out: BotCommand{
				Quote: quoteCommand{
					Securities: []bot.Security{
						{Tick: "btc", Src: enums.Satang},
					},
				},
			},
		},
		{
			In: UserCommand{
				Command: QuoteCmd,
				Chat:    "/quote bitkub btc",
			},
			Out: BotCommand{
				Quote: quoteCommand{
					Securities: []bot.Security{
						{Tick: "btc", Src: enums.Bitkub},
					},
				},
			},
		},
		{
			In: UserCommand{
				Command: TrackCmd,
				Chat:    "/track gc=f 2",
			},
			Out: BotCommand{
				Track: trackCommand{
					quoteCommand: quoteCommand{
						Securities: []bot.Security{
							{Tick: "gc=f", Src: enums.Yahoo},
						},
					},
					TrackTimes: 2,
				},
			},
		},
		{
			In: UserCommand{
				Command: TrackCmd,
				Chat:    "/track satang btc 69",
			},
			Out: BotCommand{
				Track: trackCommand{
					quoteCommand: quoteCommand{
						Securities: []bot.Security{
							{Tick: "btc", Src: enums.Satang},
						},
					},
					TrackTimes: 69,
				},
			},
		},
		{
			In: UserCommand{
				Command: TrackCmd,
				Chat:    "/track bitkub btc ada 69",
			},
			Out: BotCommand{
				Track: trackCommand{
					quoteCommand: quoteCommand{
						Securities: []bot.Security{
							{Tick: "btc", Src: enums.Bitkub},
							{Tick: "ada", Src: enums.Bitkub},
						},
					},
					TrackTimes: 69,
				},
			},
		},
		{
			In: UserCommand{
				Command: AlertCmd,
				Chat:    "/alert gc=f > 0",
			},
			Out: BotCommand{
				Alert: bot.Alert{
					Security:  bot.Security{Tick: "gc=f", Src: enums.Yahoo},
					Condition: enums.Gt,
					QuoteType: enums.Last,
					Target:    0,
				},
			},
		},
		{
			In: UserCommand{
				Command: AlertCmd,
				Chat:    "/alert gc=f bid > 0",
			},
			Out: BotCommand{
				Alert: bot.Alert{
					Security:  bot.Security{Tick: "gc=f", Src: enums.Yahoo},
					Condition: enums.Gt,
					QuoteType: enums.Bid,
					Target:    0,
				},
			},
		},
		{
			In: UserCommand{
				Command: AlertCmd,
				Chat:    "/alert satang btc bid > 112",
			},
			Out: BotCommand{
				Alert: bot.Alert{
					Security:  bot.Security{Tick: "btc", Src: enums.Satang},
					Condition: enums.Gt,
					QuoteType: enums.Bid,
					Target:    112,
				},
			},
		},
		{
			In: UserCommand{
				Command: AlertCmd,
				Chat:    "/alert bitkub btc < 112",
			},
			Out: BotCommand{
				Alert: bot.Alert{
					Security:  bot.Security{Tick: "btc", Src: enums.Bitkub},
					Condition: enums.Lt,
					QuoteType: enums.Last,
					Target:    112,
				},
			},
		},
		{
			In: UserCommand{
				Command: AlertCmd,
				Chat:    "/alert bitkub btc bid > 112",
			},
			Out: BotCommand{
				Alert: bot.Alert{
					Security:  bot.Security{Tick: "btc", Src: enums.Bitkub},
					Condition: enums.Gt,
					QuoteType: enums.Bid,
					Target:    112,
				},
			},
		},
	}

	for _, test := range tests {
		out, err := test.In.Parse()
		if err != 0 {
			t.Errorf("error parsing UserCommand: %+v\n", test.In)
		}
		switch test.In.Command {
		case QuoteCmd:
			if !reflect.DeepEqual(out.Quote.Securities, test.Out.Quote.Securities) {
				t.Errorf("invalid quote securities for: %+v\n", test.In)
			}
		case TrackCmd:
			if !reflect.DeepEqual(out.Quote.Securities, test.Out.Quote.Securities) {
				t.Errorf("invalid quote securities for: %+v\n", test.In)
			}
			if out.Track.TrackTimes != test.Out.Track.TrackTimes {
				t.Errorf("invalid track times for: %+v\n", test.In)
			}
		case AlertCmd:
			if !reflect.DeepEqual(out.Alert, test.Out.Alert) {
				if !reflect.DeepEqual(out.Alert.Security, test.Out.Alert.Security) {
					t.Errorf("invalid alert security for: %+v\n", test.In)
				}
				if out.Alert.Src != test.Out.Alert.Src {
					t.Errorf("invalid alert source for: %+v\n", test.In)
				}
				if out.Alert.Condition != test.Out.Alert.Condition {
					t.Errorf("invalid alert condition for: %+v\n", test.In)
				}
				if out.Alert.QuoteType != test.Out.Alert.QuoteType {
					t.Errorf("invalid alert quote type for: %+v\n", test.In)
				}
				if out.Alert.Target != test.Out.Alert.Target {
					t.Errorf("invalid alert target for: %+v\n", test.In)
				}
			}
		}
	}
}
