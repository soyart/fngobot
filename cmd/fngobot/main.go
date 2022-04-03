package main

import (
	"log"
	"sync"
	"time"

	tb "gopkg.in/tucnak/telebot.v3"

	"github.com/artnoi43/fngobot/cmd"
	"github.com/artnoi43/fngobot/config"
)

var (
	cmdFlags cmd.Flags
	conf     *config.Config
	tokens   []string
)

func init() {
	cmdFlags.Parse()
	log.Println("config path:", cmdFlags.ConfigFile)
	confLoc := config.ParsePath(
		cmdFlags.ConfigFile,
	)
	var err error
	conf, err = config.InitConfig(
		confLoc.Dir, confLoc.Name, confLoc.Ext,
	)
	if err != nil {
		log.Fatalf("configuration failed\n%v", err)
	}
	tokens = conf.Telegram.Client.BotTokens
}

func main() {
	var wg sync.WaitGroup
	for _, botToken := range tokens {
		b, err := tb.NewBot(tb.Settings{
			/* If empty defaults to "https://api.telegram.org" */
			URL:   "https://api.telegram.org",
			Token: botToken,
			Poller: &tb.LongPoller{
				Timeout: time.Duration(conf.Telegram.Client.TimeoutSeconds) * time.Second,
			},
			Verbose: conf.Telegram.Client.Verbose,
		})
		if err != nil {
			log.Fatalf("failed to init new bot: %s\n", err.Error())
		}

		wg.Add(1)
		go func(token string) {
			defer wg.Done()
			if err := handle(b, token); err != nil {
				log.Printf("telegram bot %s error: %s\n", token, err.Error())
			}
		}(botToken)
	}
	wg.Wait()
}
