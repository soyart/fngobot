# FnGoBot
A simple, stupid command-line-like Telegram chatbot for tracking financial market quotes written in Go

The bot is made possible with [this amazing Go Telegram bot library](https://gopkg.in/tucnak/telebot.v3). It fetches quotes from Yahoo! Finance (via [finance-go](https://github.com/piquette/finance-go)), [Binance](https://binance.com), [Coinbase](https://coinbase.com), [Satang](https://satangcorp.com), and [Bitkub](https://bitkub.com).

I try to keep the code small *and the user interface command-line-like*.

- Tested on Arch Linux and macOS.

- This bot is currently in use in my friends' circle doing real trades, so the command-line patterns are not to be changed.

# Quote sources
## `github.com/piquette/finance-go`
For Yahoo! Finance API. This is the default security quote source.

## `github.com/piquette/finance-go/crypto`
For Yahoo! Finance cryptocurrency API. This source is activated with a `crypto` switch in the command.

## `fetch/binance/binance.go`
FnGoBot's [Binance](https://binance.com) API data fetcher. This source is activated with a `binance` switch in the command.

## `fetch/coinbase/coinbase.go`
FnGoBot's [Coinbase](https://coinbase.com) API data fetcher. This source is activated with a `coinbase` switch in the command. Coinbase quotes are currently THB denominated, although this may change in the future, perhaps with adding currency configuration.

## `fetch/satang/satang.go`
FnGoBot's [satangcorp.com](https://satangcorp.com) API data fetcher. This source is activated with a `satang` switch in the command. Satang quotes are only THB denominated

## `fetch/bitkub/bitkub.go`
FnGoBot's [bitkub.com](https://bitkub.com) API data fetcher. This source is activated with a `bitkub` switch in the command. Bitkub quotes are only THB denominated.

## Running your own bot
- Clone this repository

- Provide your bot's [configuration](#config)

- Run the bot with `go run ./cmd` or build and run the binary

- If the bot successfully initializes, you may now start using the bot

- Start chatting with the bot. Send `/help` to get command-line help

## <a name="config">Configuration</a>
FnGoBot uses [Viper](github.com/spf13/viper) to manage configuration. The default config file is `$HOME/.config/fngobot/config.yml`, but `-c <config file path>` command-line option can also be used to load different configuration.
### Configuration file type
Default configuration type is YAML, but other file types like JSON can also be used.

> Only YAML and JSON are tested for FnGobot, and TOML is not supported because Telegram bot token strings contain special character `:`.

To use other configuration types (must be supported by Viper), specify the configuation file with `-c <config file>`:

    $ go run ./cmd -c path/to/config.json;

Or change the default filename in `cmd/main.go` to your desired file extension.

## External dependencies

- `gopkg.in/tucnak/telebot.v3` Telegram bot framework

- `github.com/piquette/finance-go` Yahoo! Finance data library

- `github.com/google/uuid` Google UUID library

- `github.com/spf13/viper` Configuration library

- `golang.org/x/text/message` and `golang.org/x/text/language` for thousands operator (`,`)

## Example chat commands
The chat will look like a CLI program. The bot supports 3 commands: `/quote`, `track`, and `/alert`. Source switches can be used in all commands to set quote source.

The general syntax for `/quote` is

    /quote [source switch] <ticker> [ticker..]

The example below will quote BBL.BK, KBANK.BK, KKP.BK from Yahoo! Finance:

    /quote bbl.bk kbank.bk kkp.bk

The general syntax for `/track` is

    /track [source switch] <ticker> [ticker..] <rounds>

The example below will track BTC-USD crypto pair for 2 times, each one minute apart:

    /track crypto btc-usd 2

`/alert` can handle one more switch - a bid/ask switch. This is because `/alert` can be used against last price and best bid/ask prices. If a bid/ask switch is used, it will alert based on bid or ask prices, if the bid/ask switch is not present, it will alert based on last price.

The general syntax for `/alert` is

    /alert [source switch] <ticker> [bid/ask switch] <greater/smaller> <target price>

The two modes of `/alert` are not supported all quote sources - for example, Yahoo! Finance (crypto) does not support alert with bid/ask, while Satang source does not support alert with last price.

The example below will make the bot alert if BBL.BK 'last' price is greater than 120:

    /alert bbl.bk > 120

The example below will make the bot alert if BBL.BK 'ask' price is smaller than 115:

    /alert bbl.bk ask < 115

The example below will make the bot alert if BTC-USD (from Yahoo! Finance crypto) is greater than 30,000:

    /alert crypto btc-usd > 30000

The example below will make the bot alert if BTC (From Binance) last price is smaller than 30,000:

    /alert binance btc < 30000

The example below will make the bot alert if BTC (From Coinbase) bid price is smaller than 30,000:

    /alert coinbase btc-usd bid < 30000

The example below will make the bot alert if BTC (From Bitkub.com) bid price is smaller than 30,000:

    /alert bitkub btc-usd bid < 30000

## Other chat commands

`/handlers` is used to get all running handlers in JSON format.

    /handlers

`/stop` is used to stop a running handler. Let's say we have an alerting handler whose UUID is cfd337b7, to stop it, send:

    /stop cfd337b7
