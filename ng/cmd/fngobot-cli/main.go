package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/artnoi43/fngobot/ng/adapter/handler/clihandler"
	"github.com/artnoi43/fngobot/ng/adapter/parse"
	"github.com/artnoi43/fngobot/ng/cmd"
	"github.com/artnoi43/fngobot/ng/config"
	"github.com/artnoi43/fngobot/ng/internal/enums"
)

var (
	cmdFlags cmd.Flags
	conf     *config.Config
	done     = make(chan struct{})
)

func init() {
	cmdFlags.Parse()
	log.Println("Config path:", cmdFlags.ConfigFile)
	confLoc := config.ParsePath(
		cmdFlags.ConfigFile,
	)
	var err error
	conf, err = config.InitConfig(
		confLoc.Dir, confLoc.Name, confLoc.Ext,
	)
	if err != nil {
		log.Fatalf("configuration failed\n%v", err.Error())
	}
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}
}

func main() {
	args := os.Args[1:]
	cmdStr := enums.InputCommand(args[0])
	targetBot, ok := enums.BotMap[cmdStr]
	if !ok {
		log.Fatal("invalid cmdStr")
	}

	cmd, parseError := parse.UserCommand{
		Text:      strings.Join(args, " "),
		TargetBot: targetBot,
	}.Parse()

	h := clihandler.New(&cmd, &conf.CLI, done)
	if parseError != 0 {
		h.HandleParsingError(parseError)
	} else {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			// sigChan for receiving OS signals for graceful shutdowns
			sigChan := make(chan os.Signal, 1)
			signal.Notify(
				sigChan,
				syscall.SIGHUP,  // kill -SIGHUP XXXX
				syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
				syscall.SIGQUIT, // kill -SIGQUIT XXXX
				syscall.SIGTERM, // kill -SIGTERM XXXX
			)

			// Graceful shutdown
			go func() {
				for {
					select {
					case <-sigChan:
						log.Fatal("received interrupt")
					case <-done:
						os.Exit(0)
					}
				}
			}()
		}()
		defer h.Done()
		h.Handle(targetBot)
		wg.Wait()
	}
}
