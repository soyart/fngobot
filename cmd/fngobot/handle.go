package main

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

	"github.com/artnoi43/fngobot/lib/bot/handler"
	tghandler "github.com/artnoi43/fngobot/lib/bot/handler/telegram"
	"github.com/artnoi43/fngobot/lib/bot/handler/utils"
	"github.com/artnoi43/fngobot/lib/enums"
	"github.com/artnoi43/fngobot/lib/etc/help"
	"github.com/artnoi43/fngobot/lib/parse"
)

type fngobot struct {
	history map[int64]handler.Handlers
}

func handle(b *tb.Bot, token string) error {
	f := &fngobot{
		history: make(map[int64]handler.Handlers),
	}
	log.Printf("initialized bot: %s", token)

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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-sigChan
		log.Printf("closing poller connection for %s\n", token)
		b.Close()
		log.Printf("closed poller connection for %s\n", token)
	}()

	sendFail := func() {
		log.Println("error sending Telegram message to recipient")
	}

	b.Handle("/help", handleFunc(b, f, "/help"))
	b.Handle("/quote", handleFunc(b, f, "/quote"))
	b.Handle("/track", handleFunc(b, f, "/track"))
	b.Handle("/alert", handleFunc(b, f, "/alert"))

	// Welcome/Greeting
	b.Handle("/start", func(c tb.Context) error {
		log.Println(c.Text())
		if _, err := b.Reply(c.Message(), help.LONG); err != nil {
			sendFail()
		}
		if _, err := b.Reply(c.Message(), "Hello!\nWelcome to FnGoBot chat!"); err != nil {
			sendFail()
		}
		return nil
	})

	// Stop a tracking or alerting Telegram handler
	b.Handle("/stop", func(c tb.Context) error {
		senderId := c.Sender().ID
		uuids := strings.Split(c.Text(), " ")[1:]
		for _, uuid := range uuids {
			// Stop is Handlers method
			idx, ok := f.history[senderId].Stop(uuid)
			if ok {
				// Remove handler from slice f.history[senderId]
				f.history[senderId] = append(
					f.history[senderId][:idx],
					f.history[senderId][idx+1:]...,
				)
				// Notify user for succesful removal
				if _, err := b.Reply(
					c.Message(),
					fmt.Sprintf("[%s] stopped", uuid),
				); err != nil {
					sendFail()
				}
				return nil
			}
			if _, err := b.Reply(
				c.Message(),
				fmt.Sprintf("[%s] does not exist", uuid),
			); err != nil {
				sendFail()
			}
		}
		return nil
	})

	b.Handle("/handlers", func(c tb.Context) error {
		senderId := c.Sender().ID
		handlers := f.history[senderId]
		var runningHandlers handler.Handlers
		for _, h := range handlers {
			if h.IsRunning() {
				runningHandlers = append(runningHandlers, h)
			}
		}
		var msg string
		if len(runningHandlers) > 0 {
			for _, h := range runningHandlers {
				msg = msg + utils.Yaml(h)
			}
		} else {
			msg = "No active handlers"
		}
		if _, err := b.Reply(c.Message(), msg); err != nil {
			return err
		}
		return nil
	})

	/* Catch-all help message for unhandled text */
	b.Handle(tb.OnText, func(c tb.Context) error {
		log.Println(c.Text())
		if _, err := b.Reply(
			c.Message(),
			fmt.Sprintf("wut? %s\nSee /help for help", c.Text()),
		); err != nil {
			sendFail()
		}
		return nil
	})

	go func() {
		log.Println("fngobot started")
		b.Start()
	}()

	wg.Wait()
	log.Println("fngobot exited")
	return nil
}

func handleFunc(
	b *tb.Bot,
	f *fngobot,
	command enums.InputCommand,
) func(c tb.Context) error {
	return func(c tb.Context) error {
		targetBot, exits := enums.BotMap[command]
		if !exits {
			return fmt.Errorf("invalid command")
		}
		cmd, parseError := parse.UserCommand{
			Text:      c.Text(),
			TargetBot: targetBot,
		}.Parse()
		h := tghandler.New(b, c, &cmd, conf.Telegram)
		if parseError != parse.NoErr {
			h.HandleParsingError(parseError)
			return fmt.Errorf("parseError: %d", parseError)
		}

		defer h.Done()
		senderId := c.Sender().ID
		// Append to fngobot history
		f.history[senderId] = append(f.history[senderId], h)
		h.Handle(targetBot)
		if targetBot == enums.HelpBot {
			if _, err := b.Reply(c.Message(), cmd.Help.HelpMessage); err != nil {
				return errors.Wrap(err, "failed to send help message")
			}
			return nil
		}
		return nil
	}
}
