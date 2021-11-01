package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/artnoi43/fngobot/bot/handler"
	"github.com/artnoi43/fngobot/config"
	"github.com/artnoi43/fngobot/help"
	"github.com/artnoi43/fngobot/parse"

	tb "gopkg.in/tucnak/telebot.v3"
)

// Only command-line flag is the path to configuration file
// Different bots should use different config files, not command-line flags
type flags struct {
	configFile string
}

func (f *flags) parse() {
	flag.StringVar(&f.configFile, "c", "$HOME/.config/fngobot/config.yml", "Path to configuration file")
	flag.Parse()
}

var (
	cmdFlags flags
	handlers handler.Handlers
)

func parseConfigPath() (dir, name, ext string) {
	dir, configFile := filepath.Split(cmdFlags.configFile)
	name = strings.Split(configFile, ".")[0] // remove ext from filename
	ext = filepath.Ext(configFile)[1:]       // remove dot
	return dir, name, ext
}

func init() {
	cmdFlags.parse()
	log.Println("Config path:", cmdFlags.configFile)
}

func main() {
	configPath, configFilename, configType := parseConfigPath()
	conf, err := config.InitConfig(configPath, configFilename, configType)
	if err != nil {
		log.Fatalf("Configuration failed\n%v", err)
	}

	/* Initializes bot b with TOKEN */
	b, err := tb.NewBot(tb.Settings{
		/* If empty defaults to "https://api.telegram.org" */
		URL:    "",
		Token:  conf.BotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalf("Failed to initialize bot.\nPossibly invalid token: %s", conf.BotToken)
	} else {
		log.Printf("Initialized bot: %s", conf.BotToken)
	}

	b.Handle("/help", func(c tb.Context) error {
		cmd, _ := parse.UserCommand{
			Command: parse.HelpCmd,
			Chat:    c.Text(),
		}.Parse()
		b.Send(c.Recipient(), cmd.Help.HelpMessage)
		return nil
	})

	b.Handle("/quote", func(c tb.Context) error {
		h := handler.NewHandler(b, c.Message())
		cmd, parseError := parse.UserCommand{
			Command: parse.QuoteCmd,
			Chat:    c.Text(),
		}.Parse()
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			h.SendQuote(cmd.Quote.Securities)
		}
		return nil
	})

	b.Handle("/track", func(c tb.Context) error {
		h := handler.NewHandler(b, c.Message())
		handlers = append(handlers, h)
		cmd, parseError := parse.UserCommand{
			Command: parse.TrackCmd,
			Chat:    c.Text(),
		}.Parse()
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			h.Track(cmd.Track.Securities, cmd.Track.TrackTimes, conf.BotConfig)
		}
		return nil
	})

	b.Handle("/alert", func(c tb.Context) error {
		h := handler.NewHandler(b, c.Message())
		handlers = append(handlers, h)
		cmd, parseError := parse.UserCommand{
			Command: parse.AlertCmd,
			Chat:    c.Text(),
		}.Parse()
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			h.PriceAlert(cmd.Alert, conf.BotConfig)
		}
		return nil
	})

	b.Handle("/start", func(c tb.Context) error {
		log.Println(c.Text())
		b.Send(c.Recipient(), help.LONG)
		b.Send(c.Recipient(), "Hello!\nWelcome to FnGoBot chat!")
		return nil
	})

	// Stop a tracking or alerting handler
	b.Handle("/stop", func(c tb.Context) error {
		chat := strings.Split(c.Text(), " ")
		handlers.Stop(chat[1])
		return nil
	})

	/* Catch-all help message for unhandled text */
	b.Handle(tb.OnText, func(c tb.Context) error {
		log.Println(c.Text())
		b.Send(c.Recipient(), fmt.Sprintf("wut? %s", c.Text()))
		b.Send(c.Recipient(), "Invalid argument, see /help")
		return nil
	})

	b.Start()
}
