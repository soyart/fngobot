package handler

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	tb "gopkg.in/tucnak/telebot.v3"
)

type Handler struct {
	uuid   string
	quit   chan bool
	bot    *tb.Bot
	msg    *tb.Message
}

func NewHandler(b *tb.Bot, m *tb.Message) *Handler {
	uuid := strings.Split(uuid.NewString(), "-")[0]
	quit := make(chan bool, 1)
	log.Printf("[%s]: %s (from %d)\n", uuid, m.Text, m.Sender.ID)
	return &Handler{
		uuid: uuid,
		quit: quit,
		bot:  b,
		msg:  m,
	}
}

func (h *Handler) send(s string) {
	h.bot.Send(h.msg.Sender, s)
}

func (h *Handler) notifyStop() {
	log.Printf("[%s]: Received stop signal", h.uuid)
	h.send(fmt.Sprintf("Stopping %s", h.uuid))
}