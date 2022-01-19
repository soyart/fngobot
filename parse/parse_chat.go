package parse

import (
	"strconv"
	"strings"

	"github.com/artnoi43/fngobot/bot"
	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/help"
)

type ParseError int
type Command string

const (
	// These errors are non-zero
	ErrParseInt ParseError = iota + 1
	ErrParseFloat
	ErrInvalidSign
	ErrInvalidBidAskSwitch
	ErrInvalidQuoteTypeBid
	ErrInvalidQuoteTypeAsk
	ErrInvalidQuoteTypeLast
)

const (
	HelpCmd     Command = "help"
	QuoteCmd    Command = "quote"
	TrackCmd    Command = "track"
	AlertCmd    Command = "alert"
	HandlersCmd Command = "handlers"
)

// UserCommand is essentially a chat message.
// UserCommand.Chat is an int enum
type UserCommand struct {
	Command Command
	Chat    string
}

type helpCommand struct {
	HelpMessage string `json:"help_msg"`
}

type quoteCommand struct {
	Securities []bot.Security `json:"securities"`
}
type trackCommand struct {
	quoteCommand
	TrackTimes int `json:"track_times"`
}

// BotCommand is derived from UserCommand by parsing with Parse()
// Alerting does not need its own command struct,
// as the bot.Alert struct already has all the info needed.
type BotCommand struct {
	Help  helpCommand  `json:"-"`
	Quote quoteCommand `json:"quote"`
	Track trackCommand `json:"track"`
	Alert bot.Alert    `json:"alert"`
}

// getSrc gets the source from enums,
// and also returns an index to the first ticker from the chat
func getSrc(sw string) (idx int, src enums.Src) {
	switch sw {
	case "CRYPTO":
		idx = 2
		src = enums.YahooCrypto
	case "SATANG":
		idx = 2
		src = enums.Satang
	case "BITKUB":
		idx = 2
		src = enums.Bitkub
	default:
		idx = 1
		src = enums.Yahoo
	}
	return idx, src
}

// appendSecurities receives a slice of string representing ticker
// and appends the values in the slice to cmd.Securities
func (cmd *quoteCommand) appendSecurities(ticks []string, src enums.Src) {
	for _, tick := range ticks {
		var s bot.Security
		s.Tick = tick
		s.Src = src
		cmd.Securities = append(cmd.Securities, s)
	}
}

// Parse parses UserCommand to BotCommand
func (c UserCommand) Parse() (cmd BotCommand, parseError ParseError) {
	if c.Command == HandlersCmd {
		return cmd, 0
	}
	chat := strings.Split(c.Chat, " ")
	lenChat := len(chat)
	sw := strings.ToUpper(chat[1])
	idx, src := getSrc(sw)

	switch c.Command {
	case HelpCmd:
		cmd.Help.HelpMessage = help.GetHelp(c.Chat)
	case QuoteCmd:
		cmd.Quote.appendSecurities(chat[idx:], src)
	case TrackCmd:
		cmd.Track.appendSecurities(chat[idx:lenChat-1], src)
		r, err := strconv.Atoi(chat[lenChat-1])
		if err != nil {
			return cmd, ErrParseInt
		}
		cmd.Track.TrackTimes = r
	case AlertCmd:
		cmd.Alert.Security.Tick = chat[idx]
		cmd.Alert.Security.Src = src
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

		// Determine if bid/ask or last alert
		// bid/ask alert will have length = default length + 1
		defLen := make(map[enums.Src]int)
		defLen[enums.YahooCrypto] = 5 /* /alert crypto btc-usd > 1 */
		defLen[enums.Satang] = 5      /* /alert satang btc bid > 1 */
		defLen[enums.Bitkub] = 5      /* /alert bitkub btc > 1 */
		defLen[enums.Yahoo] = 4       /* /alert bbl.bk > 120 */
		// If unsupported alert/ use ones that are supported
		switch lenChat {
		// Last price alerts
		case defLen[cmd.Alert.Src]:
			// Satang does not support last price
			if cmd.Alert.Src == enums.Satang {
				return cmd, ErrInvalidQuoteTypeLast
			}
			cmd.Alert.QuoteType = enums.Last
		// Bid/ask alerts
		default:
			bidask := strings.ToUpper(chat[idx+1])
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
	return cmd, parseError
}
