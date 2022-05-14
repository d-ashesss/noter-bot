package app

import (
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"strings"
)

func initBotHandlers(b *telebot.Bot, a *App) {
	b.Handle(telebot.OnText, a.botHandleTextMessage)
}

func (a App) botHandleTextMessage(m *telebot.Message) {
	log.Printf("[bot] incoming message: %s: %s", getTelebotUserName(m.Sender), m.Text)
}

func getTelebotUserName(user *telebot.User) string {
	if len(user.Username) > 0 {
		return "@" + user.Username
	}
	return strings.Trim(user.FirstName+" "+user.LastName, " ")
}
