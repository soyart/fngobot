package main

import (
	"flag"
	"log"
	"os"
	"strings"

	clihandler "github.com/artnoi43/fngobot/bot/handler_cli"
	"github.com/artnoi43/fngobot/config"
	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/parse"
)

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

func init() {
	cmdFlags.parse()
}

func main() {
	confPath, confFile, confType := config.ParseConfigPath(cmdFlags.configFile)
	conf, err := config.InitConfig(confPath, confFile, confType)
	if err != nil {
		log.Fatalf("failed to init config: %v\n", err.Error())
	}

	args := os.Args[1:]
	cmdStr := enums.Command(args[0])
	targetBot := enums.CommandMap[cmdStr]

	cmd, parseError := parse.UserCommand{
		Type: targetBot,
		Text: strings.Join(args, " "),
	}.Parse()

	h := clihandler.New(&cmd, conf)
	if parseError != 0 {
		h.HandleParsingError(parseError)
	} else {
		h.Handle(targetBot)
	}
}
