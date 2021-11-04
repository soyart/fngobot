package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/artnoi43/fngobot/bot/handler"
	"github.com/artnoi43/fngobot/config"
	"github.com/artnoi43/fngobot/help"
	"github.com/artnoi43/fngobot/parse"

	tb "gopkg.in/tucnak/telebot.v2"
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

	b.Handle("/help", func(m *tb.Message) {
		cmd, _ := parse.UserCommand{
			Command: parse.HelpCmd,
			Chat:    m.Text,
		}.Parse()
		b.Send(m.Sender, cmd.Help.HelpMessage)
	})

	b.Handle("/quote", func(m *tb.Message) {
		cmd, parseError := parse.UserCommand{
			Command: parse.QuoteCmd,
			Chat:    m.Text,
		}.Parse()
		h := handler.NewHandler(b, m, conf.BotConfig, &cmd)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			h.Handle(handler.QUOTEBOT)
		}
	})

	b.Handle("/track", func(m *tb.Message) {
		cmd, parseError := parse.UserCommand{
			Command: parse.TrackCmd,
			Chat:    m.Text,
		}.Parse()
		h := handler.NewHandler(b, m, conf.BotConfig, &cmd)
		handlers = append(handlers, h)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			h.Handle(handler.TRACKBOT)
		}
	})

	b.Handle("/alert", func(m *tb.Message) {
		cmd, parseError := parse.UserCommand{
			Command: parse.AlertCmd,
			Chat:    m.Text,
		}.Parse()
		h := handler.NewHandler(b, m, conf.BotConfig, &cmd)
		handlers = append(handlers, h)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			h.Handle(handler.ALERTBOT)
		}
	})

	b.Handle("/start", func(m *tb.Message) {
		log.Println(m.Text)
		b.Send(m.Sender, help.LONG)
		b.Send(m.Sender, "Hello!\nWelcome to FnGoBot chat!")
	})

	// Stop a tracking or alerting handler
	b.Handle("/stop", func(m *tb.Message) {
		uuid := strings.Split(m.Text, " ")[1]
		i, ok := handlers.Stop(uuid)
		if ok {
			handlers = append(handlers[:i], handlers[i+1:]...)
		}
	})

	/* Catch-all help message for unhandled text */
	b.Handle(tb.OnText, func(m *tb.Message) {
		log.Println(m.Text)
		b.Send(m.Sender, "wut? "+m.Text)
		b.Send(m.Sender, "Invalid argument, see /help")
	})

	b.Start()
}
