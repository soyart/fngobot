package help

import (
	"strings"
)

func GetHelp(s string) string {
	chat := strings.Split(s, " ")
	var messg string

	if len(chat) == 1 {
		messg = SHORT
	} else {
		switch strings.ToLower(chat[1]) {
		case "stop":
			messg = STOP
		case "quote":
			messg = QUOTE
		case "track":
			messg = TRACK
		case "alert":
			messg = ALERT
		case "long", "all", "-l", "-a":
			messg = LONG + QUOTE + TRACK + ALERT + HandlersBot
		default:
			messg = SHORT
		}
	}
	return messg
}

const (
	SHORT = `Type '/help long' to see long-ass help message for all available commands

Alternatively, use /help <command>:

/help quote
for /quote command 

/help track
for /track command

/help alert
for /alert command.

/help all
for all commands`

	LONG = `@FnGoBot is a bot for quoting, tracking, and alerting financial market security prices. By default, the bot uses Yahoo! Finance ticker symbols, visit this link for reference: https://finance.yahoo.com/lookup

The bot code is released as free software (BSD licensed), and its source code is available on https://github.com/artnoi43/fngobot

The bot currently handles 3 commands:
(1) /quote - quote price
(2) /track - track price
(3) /alert - alert when conditions met

The command can be used with one of the switches 'crypto', 'satang', and 'bitkub', i.e. use '/quote bitkub <tick>' will get recent quotes from 'tick' on Bitkub exchange. If no switch is used, the bot gets its quotes from Yahoo! Finance.

This bot categorizes securities into 2 types: (1) non-cryptocurrencies and (2) cryptocurrencies. Both types are handled by the same commands.

Cryptocurrencies need 'crypto' or 'satang' keyword after the command to tell the bot that the following tickers are cryptocurrencies.

If multiple tickers are used, cryptocurrencies and non-crypto securities can NOT be mixed in the same command.

`

	STOP = `/stop - stop a tracking/alerting handler

Syntax:

/stop <handler UUID>
`

	QUOTE = `/quote - get bid/ask or market quotes

Syntax:

/quote [switch] <tick0> [tick1..tickN]

Examples:

/quote OR.BK

will get current quotes on OR.BK

/quote OR.BK BBL.BK FOO.BAR

will get current quotes on OR.BK, BBL.BK and FOO.BAR

/quote crypto BTC-USD DOGE-USD

will get current values of BTC-USD and DOGE-USD

/quote satang BTC EOS ZEC

will get current quote on BTC-THB, EOS-THB, and ZEC-THB on Satang exchange

/quote bitkub ADA

will get current quote on ADA-THB on Bitkub exchange`

	TRACK = `/track - periodically track prices for user-specified timeframe. /track will track commands for n minutes, once every minute.

Syntax:
/track [switch] <tick0> [tick1..tickN] <minutes>

Examples:

/track OR.BK 3

will track OR.BK quotes from Yahoo! Finance for 3 minutes

/track OR.BK BBL.BK 4

will track OR.BK and BBL.BK quotes from Yahoo! Finance for 4 minutes

/track crypto DOGE-USD 15

will track cryptocurrency DOGE-USD pair from Yahoo! Finance for 3 minutes

/track satang DOGE 2

will track cryptocurrency DOGE-THB pair on Satang exchange for 2 minutes

/track bitkub DOGE BCH 7

will track cryptocurrency DOGE-THB and BCH-THB pairs on Bitkub exchange for 7 minutes`

	ALERT = `/alert - like track, but the bot will ONLY send messages when the prices match user-supplied condition.
Condition is supplied as '>' or '<' (>= or <=)
Bid/ask switch can also be used on certain securities

Syntax:
/alert [switch] <tick> [bid/ask] <condition> <target price>

/alert can NOT be used with multiple tickers

For regular stocks and Bitkub, you can omit bid/ask switch:

/alert OR.BK > 30
/alert bitkub xrp > 30

Will alert if last price is greater than 30

Or if you choose to supply bid/ask switch:

/alert OR.BK bid > 30
/alert bitkub bid xrp > 30

Will alert if best bid is greater than 30

Unfortunately, Crypto do NOT support bid/ask switch:

/alert crypto btc-usd > 40000
/alert crypto eth-usd < 3000

Unfortunately, Satang quotes ONLY support bid/ask

/alert satang ada ask < 60
/alert satang btc bid > 1400000`

	HandlersBot = `/handlers - get active handlers in JSON
This command does not need extra chat message,
i.e. you can just send the bot '/handlers'`
)
