package main

import (
	"log"
	"os"
	"strings"

	clihandler "github.com/artnoi43/fngobot/bot/handler/cli"
	"github.com/artnoi43/fngobot/cmd"
	"github.com/artnoi43/fngobot/config"
	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/parse"
)

var (
	cmdFlags cmd.Flags
	conf     *config.Config
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
	cmdStr := enums.Command(args[0])
	targetBot, ok := enums.BotMap[cmdStr]
	if !ok {
		log.Fatal("invalid cmdStr")
	}

	cmd, parseError := parse.UserCommand{
		Type: targetBot,
		Text: strings.Join(args, " "),
	}.Parse()

	h := clihandler.New(&cmd, &conf.CLI)
	if parseError != 0 {
		h.HandleParsingError(parseError)
	} else {
		defer h.Done()
		h.Handle(targetBot)
	}
}
