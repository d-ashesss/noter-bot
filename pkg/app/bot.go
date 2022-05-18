package app

import (
	"context"
	"github.com/d-ashesss/noter-bot/pkg/model"
	"gopkg.in/tucnak/telebot.v2"
	"log"
)

func initBotHandlers(b *telebot.Bot, a *App) {
	b.Handle(telebot.OnText, a.botHandleTextMessage)
}

func (a App) botHandleTextMessage(m *telebot.Message) {
	n := model.NewNote(m.Sender.ID, m.Text)
	if err := a.noteModel.Create(context.TODO(), n); err != nil {
		log.Printf("[bot] failed to save note: %s", err)

		if _, err := a.bot.Reply(m, "Failed to save this note 😥"); err != nil {
			log.Printf("[bot] > failed to notify user: %s", err)
		}
		return
	}
	log.Printf("[bot] created new note %s for user %d", n.ID, n.UserID)
	_, err := a.bot.Send(
		m.Sender,
		n.Text,
		&telebot.SendOptions{ParseMode: telebot.ModeMarkdown},
	)
	if err != nil {
		log.Printf("[bot] failed to display note %s: %s", n.ID, err)
		return
	}
	if err := a.bot.Delete(m); err != nil {
		log.Printf("[bot] failed to delete original for note %s: %s", n.ID, err)
	}
}
