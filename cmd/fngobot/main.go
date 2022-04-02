package main

import (
	"log"
	"sync"

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
	tokens = conf.Telegram.BotTokens
}

func main() {
	var wg sync.WaitGroup
	for _, botToken := range tokens {
		wg.Add(1)
		go func(token string) {
			defer wg.Done()
			if err := start(token); err != nil {
				log.Println("telegram bot error", err)
			}
		}(botToken)
	}
	wg.Wait()
}
