package parse

import (
	"log"
	"strconv"
	"strings"

	"github.com/artnoi43/mgl/str"

	"github.com/artnoi43/fngobot/adapter/fetch"
	"github.com/artnoi43/fngobot/internal/enums"
	"github.com/artnoi43/fngobot/internal/help"
	"github.com/artnoi43/fngobot/usecase"
)

type UserCommand struct {
	Text      string        `json:"chat,omitempty"`
	TargetBot enums.BotType `json:"command,omitempty"`
}

type helpCommand struct {
	HelpMessage string `json:"help_msg,omitempty"`
}

type quoteCommand struct {
	Securities []usecase.Security `json:"securities,omitempty"`
}
type trackCommand struct {
	quoteCommand
	TrackTimes int `json:"track_times,omitempty"`
}

// BotCommand is derived from UserCommand by parsing with Parse()
// Alerting does not need its own command struct,
// as the usecase.Alert struct already has all the info needed.
type BotCommand struct {
	Help  helpCommand   `json:"-"`
	Quote quoteCommand  `json:"quote,omitempty"`
	Track trackCommand  `json:"track,omitempty"`
	Alert usecase.Alert `json:"alert,omitempty"`
}

// getSrc returns the source and from enums.SwitchSrcMap,
// and if it's not available, returns 1, enums.Yahoo
func getSrc(sw string) (src enums.Src, startIdx int) {
	target, exists := enums.SwitchSrcMap[enums.Switch(str.ToUpper(sw))]
	if !exists {
		return enums.Yahoo, 1
	}
	if len(target) == 1 {
		for src, startIdx := range target {
			return src, startIdx
		}
	}
	return enums.Yahoo, 1
}

// appendSecurities receives a slice of string representing ticker
// and appends the values in the slice to cmd.Securities
func (cmd *quoteCommand) appendSecurities(ticks []string, src enums.Src) {
	for _, tick := range ticks {
		var s usecase.Security
		s.Fetcher = fetch.New(src)
		s.Tick = str.ToUpper(tick)
		s.Src = src
		cmd.Securities = append(cmd.Securities, s)
	}
}

// Parse parses UserCommand to BotCommand
func (c UserCommand) Parse() (cmd BotCommand, e ParseError) {
	targetBot := c.TargetBot
	if targetBot == enums.HandlersBot {
		return cmd, 0
	}
	if !targetBot.IsValid() {
		log.Fatalf("parse: invalid bot type: %s\n", targetBot)
	}
	chat := strings.Fields(c.Text)
	lenChat := len(chat)
	var sw string
	if targetBot != enums.HelpBot {
		sw = str.ToUpper(chat[1])
	}
	src, startIdx := getSrc(sw)
	switch targetBot {
	case enums.HelpBot:
		cmd.Help.HelpMessage = help.GetHelp(c.Text)
	case enums.QuoteBot:
		cmd.Quote.appendSecurities(chat[startIdx:], src)
	case enums.TrackBot:
		cmd.Track.appendSecurities(chat[startIdx:lenChat-1], src)
		r, err := strconv.Atoi(chat[lenChat-1])
		if err != nil {
			return cmd, ErrParseInt
		}
		cmd.Track.TrackTimes = r
	case enums.AlertBot:
		cmd.Alert.Security.Tick = str.ToUpper(chat[startIdx])
		cmd.Alert.Security.Src = src
		cmd.Alert.Security.Fetcher = fetch.New(src)
		targ, err := strconv.ParseFloat(chat[lenChat-1], 64)
		if err != nil {
			return cmd, ErrParseFloat
		}
		cmd.Alert.Target = targ

		sign := chat[lenChat-2]
		if sign != "<" && sign != ">" {
			return cmd, ErrInvalidSign
		} else if sign == "<" {
			cmd.Alert.Condition = enums.Lt
		} else if sign == ">" {
			cmd.Alert.Condition = enums.Gt
		}

		// If unsupported alert/ use ones that are supported
		switch lenChat {
		// Last price alerts
		case startIdx + 3:
			// Satang does not support last price
			if cmd.Alert.Src == enums.Satang {
				return cmd, ErrInvalidQuoteTypeLast
			}
			cmd.Alert.QuoteType = enums.Last
			// Bid/ask alerts
		default:
			bidask := str.ToUpper(chat[startIdx+1])
			switch cmd.Alert.Src {
			// Yahoo Crypto does not support bid/ask price
			case enums.YahooCrypto:
				if bidask == "BID" {
					return cmd, ErrInvalidQuoteTypeBid
				} else if bidask == "ASK" {
					return cmd, ErrInvalidQuoteTypeAsk
				}
			default:
				if bidask == "BID" {
					cmd.Alert.QuoteType = enums.Bid
				} else if bidask == "ASK" {
					cmd.Alert.QuoteType = enums.Ask
				} else {
					return cmd, ErrInvalidBidAskSwitch
				}
			}
		}
	}
	return cmd, e
}
