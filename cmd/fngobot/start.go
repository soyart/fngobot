package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	tb "gopkg.in/tucnak/telebot.v3"

	tghandler "github.com/artnoi43/fngobot/bot/handler/telegram"
	"github.com/artnoi43/fngobot/enums"
	"github.com/artnoi43/fngobot/help"
	"github.com/artnoi43/fngobot/parse"
	"github.com/pkg/errors"
)

func start(token string) error {
	b, err := tb.NewBot(tb.Settings{
		/* If empty defaults to "https://api.telegram.org" */
		URL:    "",
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return errors.Wrapf(err, "failed to initialize bot.\nPossibly invalid token: %s", token)
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
		log.Println("closed poller connection")
	}()

	sendFail := func() {
		log.Println("error sending Telegram message to recipient")
	}

	b.Handle("/help", func(c tb.Context) error {
		cmd, _ := parse.UserCommand{
			Type: enums.HelpBot,
			Text: c.Text(),
		}.Parse()
		if _, err := b.Reply(c.Message(), cmd.Help.HelpMessage); err != nil {
			sendFail()
		}
		return nil
	})

	b.Handle("/quote", func(c tb.Context) error {
		cmd, parseError := parse.UserCommand{
			Type: enums.QuoteBot,
			Text: c.Text(),
		}.Parse()
		h := tghandler.New(b, c, &cmd, conf.Telegram)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			defer h.Done()
			h.Handle(enums.QuoteBot)
		}
		return nil
	})

	b.Handle("/track", func(c tb.Context) error {
		cmd, parseError := parse.UserCommand{
			Type: enums.TrackBot,
			Text: c.Text(),
		}.Parse()
		h := tghandler.New(b, c, &cmd, conf.Telegram)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			defer h.Done()
			h.Handle(enums.TrackBot)
		}
		return nil
	})

	b.Handle("/alert", func(c tb.Context) error {
		cmd, parseError := parse.UserCommand{
			Type: enums.AlertBot,
			Text: c.Text(),
		}.Parse()
		h := tghandler.New(b, c, &cmd, conf.Telegram)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			defer h.Done()
			h.Handle(enums.AlertBot)
		}
		return nil
	})

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

	// Stop a tracking or alerting tghandler
	b.Handle("/handlers", func(c tb.Context) error {
		cmd, parseError := parse.UserCommand{
			Type: enums.HandlersBot,
			Text: c.Text(),
		}.Parse()
		h := tghandler.New(b, c, &cmd, conf.Telegram)
		if parseError != 0 {
			h.HandleParsingError(parseError)
		} else {
			defer h.Done()
			h.Handle(enums.HandlersBot)
		}
		return nil
	})

	// Stop a tracking or alerting tghandler
	b.Handle("/stop", func(c tb.Context) error {
		senderId := c.Sender().ID
		uuids := strings.Split(c.Text(), " ")[1:]
		for _, uuid := range uuids {
			// Stop is Handlers method
			idx, ok := tghandler.SenderHandlers[senderId].Stop(uuid)
			if ok {
				// Remove is a plain function
				tghandler.Remove(senderId, idx)
			}
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
