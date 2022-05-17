package app

import (
	"github.com/d-ashesss/noter-bot/pkg/model"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"strings"
)

func initBotHandlers(b *telebot.Bot, a *App) {
	b.Handle(telebot.OnText, a.botHandleTextMessage)
}

func (a App) botHandleTextMessage(m *telebot.Message) {
	log.Printf("[bot] incoming message: %s: %s", getTelebotUserName(m.Sender), m.Text)
	n := model.NewNote(m.Sender.ID, m.Text)
	log.Printf("[bot] created new note %v", n)
	_, err := a.bot.Send(m.Sender, n.Text)
	if err != nil {
		log.Printf("[bot] note response: %s", err)
	} else {
		_ = a.bot.Delete(m)
	}
}

func getTelebotUserName(user *telebot.User) string {
	if len(user.Username) > 0 {
		return "@" + user.Username
	}
	return strings.Trim(user.FirstName+" "+user.LastName, " ")
}
