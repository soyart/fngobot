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

	/* Initializes bot b with token conf.BotToken */
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
		cmd, parseError := parse.UserCommand{
			Command: parse.QuoteCmd,
			Chat:    c.Text(),
		}.Parse()
		h := handler.NewHandler(b, c.Message(), conf.BotConfig, &cmd)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			defer h.Done()
			h.Handle(handler.QUOTEBOT)
		}
		return nil
	})

	b.Handle("/track", func(c tb.Context) error {
		cmd, parseError := parse.UserCommand{
			Command: parse.TrackCmd,
			Chat:    c.Text(),
		}.Parse()
		h := handler.NewHandler(b, c.Message(), conf.BotConfig, &cmd)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			defer h.Done()
			h.Handle(handler.TRACKBOT)
		}
		return nil
	})

	b.Handle("/alert", func(c tb.Context) error {
		cmd, parseError := parse.UserCommand{
			Command: parse.AlertCmd,
			Chat:    c.Text(),
		}.Parse()
		h := handler.NewHandler(b, c.Message(), conf.BotConfig, &cmd)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			defer h.Done()
			h.Handle(handler.ALERTBOT)
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
	b.Handle("/handlers", func(c tb.Context) error {
		cmd, parseError := parse.UserCommand{
			Command: parse.HandlersCmd,
			Chat:    c.Text(),
		}.Parse()
		h := handler.NewHandler(b, c.Message(), conf.BotConfig, &cmd)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			defer h.Done()
			h.Handle(handler.HANDLERS)
		}
		return nil
	})

	// Stop a tracking or alerting handler
	b.Handle("/stop", func(c tb.Context) error {
		uuids := strings.Split(c.Text(), " ")[1:]
		for _, uuid := range uuids {
			i, ok := handler.BotHandlers.Stop(uuid)
			if ok {
				handler.BotHandlers = append(handler.BotHandlers[:i], handler.BotHandlers[i+1:]...)
			}
		}
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
