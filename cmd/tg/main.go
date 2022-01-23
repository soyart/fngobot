package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/artnoi43/fngobot/bot/tghandler"
	"github.com/artnoi43/fngobot/config"
	"github.com/artnoi43/fngobot/enums"
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

	token := conf.Telegram.BotToken
	/* Initializes bot b with token */
	b, err := tb.NewBot(tb.Settings{
		/* If empty defaults to "https://api.telegram.org" */
		URL:    "",
		Token:  conf.Telegram.BotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	var msg string
	if err != nil {
		msg = "Failed to initialize bot.\nPossibly invalid token: %s"
	} else {
		msg = "Initialized bot: %s"
	}
	log.Printf(msg, token)

	b.Handle("/help", func(c tb.Context) error {
		cmd, _ := parse.UserCommand{
			Command: parse.HelpCmd,
			Chat:    c.Text(),
		}.Parse()
		b.Send(c.Recipient(), cmd.Help.HelpMessage)
		return nil
	})

	b.Handle("/quote", func(c tb.Context) error {
		cmd, parseError := parse.UserCommand{
			Command: parse.QuoteCmd,
			Chat:    c.Text(),
		}.Parse()
		h := tghandler.New(b, c.Message(), conf.Telegram, &cmd)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			defer h.Done()
			h.Handle(enums.QUOTEBOT)
		}
		return nil
	})

	b.Handle("/track", func(c tb.Context) error {
		cmd, parseError := parse.UserCommand{
			Command: parse.TrackCmd,
			Chat:    c.Text(),
		}.Parse()
		h := tghandler.New(b, c.Message(), conf.Telegram, &cmd)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			defer h.Done()
			h.Handle(enums.TRACKBOT)
		}
		return nil
	})

	b.Handle("/alert", func(c tb.Context) error {
		cmd, parseError := parse.UserCommand{
			Command: parse.AlertCmd,
			Chat:    c.Text(),
		}.Parse()
		h := tghandler.New(b, c.Message(), conf.Telegram, &cmd)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			defer h.Done()
			h.Handle(enums.ALERTBOT)
		}
		return nil
	})

	b.Handle("/start", func(c tb.Context) error {
		log.Println(c.Text())
		b.Send(c.Recipient(), help.LONG)
		b.Send(c.Recipient(), "Hello!\nWelcome to FnGoBot chat!")
		return nil
	})

	// Stop a tracking or alerting tghandler
	b.Handle("/handlers", func(c tb.Context) error {
		cmd, parseError := parse.UserCommand{
			Command: parse.HandlersCmd,
			Chat:    c.Text(),
		}.Parse()
		h := tghandler.New(b, c.Message(), conf.Telegram, &cmd)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			defer h.Done()
			h.Handle(enums.HANDLERS)
		}
		return nil
	})

	// Stop a tracking or alerting tghandler
	b.Handle("/stop", func(c tb.Context) error {
		senderId := c.Sender().ID
		uuids := strings.Split(c.Text(), " ")[1:]
		for _, uuid := range uuids {
			// Stop is Handlers method
			idx, ok := tghandler.SenderHandlers[senderId].Stop(uuid)
			if ok {
				// Remove is a plain function
				tghandler.Remove(senderId, idx)
			}
		}
		return nil
	})

	/* Catch-all help message for unhandled text */
	b.Handle(tb.OnText, func(c tb.Context) error {
		log.Println(c.Text())
		b.Send(
			c.Recipient(),
			fmt.Sprintf("wut? %s\nSee /help for help", c.Text()),
		)
		return nil
	})

	b.Start()
}
