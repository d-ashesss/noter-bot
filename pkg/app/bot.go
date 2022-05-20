package app

import (
	"github.com/d-ashesss/noter-bot/pkg/model"
	"gopkg.in/tucnak/telebot.v2"
	"log"
)

const (
	botCmdStart = "/start"

	botMessageWelcome = "*Welcome!*\n" +
		"With this bot you will be able to save and manage your personal notes.\n" +
		"To create a note simply send me a text message."
	botMessageFailedToSave = "Failed to save this note 😥"
)

func initBotHandlers(b *telebot.Bot, a *App) {
	b.Handle(botCmdStart, a.botHandleStartCommand)
	b.Handle(telebot.OnText, a.botHandleTextMessage)
}

func (a App) botHandleStartCommand(m *telebot.Message) {
	if _, err := a.bot.Send(
		m.Sender,
		botMessageWelcome,
		&telebot.SendOptions{ParseMode: telebot.ModeMarkdown},
	); err != nil {
		log.Printf("[bot] failed to welcome user: %s", err)
	}
}

func (a App) botHandleTextMessage(m *telebot.Message) {
	n := model.NewNote(m.Sender.ID, m.Text)
	if err := a.noteModel.Create(a.botCtx, n); err != nil {
		log.Printf("[bot] failed to save note: %s", err)

		if _, err := a.bot.Reply(m, botMessageFailedToSave); err != nil {
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
