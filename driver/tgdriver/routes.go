package tgdriver

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/pkg/errors"
	tb "gopkg.in/tucnak/telebot.v3"

	"github.com/artnoi43/fngobot/adapter/handler"
	"github.com/artnoi43/fngobot/adapter/handler/tghandler"
	"github.com/artnoi43/fngobot/adapter/handler/utils"
	"github.com/artnoi43/fngobot/adapter/parse"
	"github.com/artnoi43/fngobot/internal/enums"
	"github.com/artnoi43/fngobot/internal/help"
)

type fngobot struct {
	// Telegram's SenderID is int64
	history map[int64]handler.Handlers
}

// NewBot inits a new bot, registers its routes, and start running the bot
func (tg *tgDriver) InitAndStartBot() error {
	fngobot := &fngobot{
		history: make(map[int64]handler.Handlers),
	}
	log.Printf("initialized bot: %s\n", tg.token)

	// sigChan for receiving OS signals for graceful shutdowns
	sigChan := make(chan os.Signal, 1)
	signal.Notify(
		sigChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+ctx
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
		syscall.SIGTERM, // kill -SIGTERM XXXX
	)

	// Graceful shutdown
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-sigChan
		log.Printf("closing poller connection for %s\n", tg.token)
		tg.bot.Close()
		log.Printf("closed poller connection for %s\n", tg.token)
	}()

	sendFail := func() {
		log.Println("error sending Telegram message to recipient")
	}

	tg.bot.Handle("/help", tg.handleFunc(fngobot, "/help"))
	tg.bot.Handle("/quote", tg.handleFunc(fngobot, "/quote"))
	tg.bot.Handle("/track", tg.handleFunc(fngobot, "/track"))
	tg.bot.Handle("/alert", tg.handleFunc(fngobot, "/alert"))

	// Welcome/Greeting
	tg.bot.Handle("/start", func(ctx tb.Context) error {
		log.Println(ctx.Text())
		if _, err := tg.bot.Reply(ctx.Message(), help.LONG); err != nil {
			sendFail()
		}
		if _, err := tg.bot.Reply(ctx.Message(), "Hello!\nWelcome to FnGoBot chat!"); err != nil {
			sendFail()
		}
		return nil
	})

	// Stop a tracking or alerting Telegram handler
	tg.bot.Handle("/stop", func(ctx tb.Context) error {
		senderId := ctx.Sender().ID
		uuids := strings.Split(ctx.Text(), " ")[1:]
		for _, uuid := range uuids {
			// Stop is Handlers method
			idx, ok := fngobot.history[senderId].Stop(uuid)
			if ok {
				// Remove handler from slice fngobot.history[senderId]
				fngobot.history[senderId] = append(
					fngobot.history[senderId][:idx],
					fngobot.history[senderId][idx+1:]...,
				)
				// Notify user for succesful removal
				if _, err := tg.bot.Reply(
					ctx.Message(),
					fmt.Sprintf("[%s] stopped", uuid),
				); err != nil {
					sendFail()
				}
				return nil
			}
			if _, err := tg.bot.Reply(
				ctx.Message(),
				fmt.Sprintf("[%s] does not exist", uuid),
			); err != nil {
				sendFail()
			}
		}
		return nil
	})

	tg.bot.Handle("/handlers", func(ctx tb.Context) error {
		senderId := ctx.Sender().ID
		var runningHandlers handler.Handlers
		var msg string
		// findAndFormat finds runningHandlers and formats msg.
		// I made it a func in case we want to reuse it in this block.
		findAndFormat := func() {
			// Overwrites runningHandlers and msg
			runningHandlers = handler.Handlers{}
			msg = ""
			handlers := fngobot.history[senderId]
			for _, h := range handlers {
				if h.IsRunning() {
					runningHandlers = append(runningHandlers, h)
				}
			}
			if len(runningHandlers) > 0 {
				for _, h := range runningHandlers {
					msg = msg + utils.Yaml(h)
				}
			} else {
				msg = "No active handlers"
			}
		}
		findAndFormat()
		if _, err := tg.bot.Reply(ctx.Message(), msg); err != nil {
			return err
		}

		return nil
	})

	/* Catch-all help message for unhandled text */
	tg.bot.Handle(tb.OnText, func(ctx tb.Context) error {
		log.Println(ctx.Text())
		if _, err := tg.bot.Reply(
			ctx.Message(),
			fmt.Sprintf("wut? %s\nSee /help for help", ctx.Text()),
		); err != nil {
			sendFail()
		}
		return nil
	})

	go func() {
		log.Println("fngobot started")
		tg.bot.Start()
	}()

	wg.Wait()
	log.Println("fngobot exited")
	return nil
}

func (tg *tgDriver) handleFunc(
	fngobot *fngobot,
	command enums.InputCommand,
) func(ctx tb.Context) error {
	return func(ctx tb.Context) error {
		targetBot, exists := enums.BotMap[command]
		if !exists {
			return fmt.Errorf("invalid command '%s'", command)
		}
		cmd, parseError := parse.UserCommand{
			Text:      ctx.Text(),
			TargetBot: targetBot,
		}.Parse()

		// New *tg.handler
		h := tghandler.New(tg.bot, ctx, &cmd, *tg.handlerConfig)
		if parseError != parse.NoErr {
			h.HandleParsingError(parseError)
			return fmt.Errorf("parseError: %d", parseError)
		}

		defer h.Done()
		senderId := ctx.Sender().ID
		// Append to fngobot history
		fngobot.history[senderId] = append(fngobot.history[senderId], h)
		h.Handle(targetBot)
		if targetBot == enums.HelpBot {
			if _, err := tg.bot.Reply(ctx.Message(), cmd.Help.HelpMessage); err != nil {
				return errors.Wrap(err, "failed to send help message")
			}
		}
		return nil
	}
}
