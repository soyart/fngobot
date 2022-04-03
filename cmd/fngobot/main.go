package main

import (
	"log"
	"sync"
	"time"

	"github.com/artnoi43/fngobot/cmd"
	"github.com/artnoi43/fngobot/config"
	tb "gopkg.in/tucnak/telebot.v3"
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
	tokens = conf.Telegram.BotTokens
}

func main() {
	var wg sync.WaitGroup
	for _, botToken := range tokens {
		b, err := tb.NewBot(tb.Settings{
			/* If empty defaults to "https://api.telegram.org" */
			URL:    "",
			Token:  botToken,
			Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("failed to init new bot: %s\n", err.Error())
		}

		wg.Add(1)
		go func(token string) {
			defer wg.Done()
			if err := runBot(b, token); err != nil {
				log.Println("telegram bot error", err)
			}
		}(botToken)
	}
	wg.Wait()
}
