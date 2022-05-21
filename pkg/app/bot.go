package app

import (
	"github.com/d-ashesss/noter-bot/pkg/model"
	"gopkg.in/tucnak/telebot.v2"
	"log"
)

const (
	BotCommandStart = "/start"

	BotMessageWelcome = "*Welcome!*\n" +
		"With this bot you will be able to save and manage your personal notes.\n" +
		"To create a note simply send me a text message."
	BotMessageNoteFailedToSave   = "😥 Failed to save this note."
	BotMessageNoteFailedToDelete = "❗️ Failed to delete the note."
	BotMessageNoteDeleted        = "✅ Note deleted."
)

func initBotHandlers(b *telebot.Bot, a *App) {
	b.Handle(BotCommandStart, a.botHandleCommandStart)
	b.Handle(telebot.OnText, a.botHandleMessageText)

	noteOptionsMenu := NewBotMenuNoteOptions("")
	b.Handle(noteOptionsMenu.BtnDelete, a.botHandleCallbackNoteOptionsDelete)
}

func (a *App) botHandleCallbackNoteOptionsDelete(cb *telebot.Callback) {
	log.Printf("[bot] request to delete note %q", cb.Data)
	n, err := a.noteModel.Get(a.botCtx, cb.Data)
	if err == nil {
		if err := a.noteModel.Delete(a.botCtx, n); err != nil {
			log.Printf("[bot] failed to delete note %s: %s", n.ID, err)
			_ = a.bot.Respond(cb, &telebot.CallbackResponse{Text: BotMessageNoteFailedToDelete, ShowAlert: true})
			return
		}
		log.Printf("[bot] note %s was deleted", n.ID)
	} else {
		log.Printf("[bot] note %q was not found: %s", cb.Data, err)
	}
	_ = a.bot.Respond(cb, &telebot.CallbackResponse{Text: BotMessageNoteDeleted})
	if err := a.bot.Delete(cb.Message); err != nil {
		log.Printf("[bot] failed to delete message with note %s: %s", n.ID, err)
	}
}

func (a *App) botHandleCommandStart(m *telebot.Message) {
	if _, err := a.bot.Send(
		m.Sender,
		BotMessageWelcome,
		&telebot.SendOptions{ParseMode: telebot.ModeMarkdown},
	); err != nil {
		log.Printf("[bot] failed to welcome user: %s", err)
	}
}

func (a *App) botHandleMessageText(m *telebot.Message) {
	n := model.NewNote(m.Sender.ID, m.Text)
	if err := a.noteModel.Create(a.botCtx, n); err != nil {
		log.Printf("[bot] failed to save note: %s", err)

		if _, err := a.bot.Reply(m, BotMessageNoteFailedToSave); err != nil {
			log.Printf("[bot] > failed to notify user: %s", err)
		}
		return
	}
	log.Printf("[bot] created new note %s for user %d", n.ID, n.UserID)
	_, err := a.bot.Send(
		m.Sender,
		n.Text,
		&telebot.SendOptions{ParseMode: telebot.ModeMarkdown},
		NewBotMenuNoteOptions(n.ID).Menu,
	)
	if err != nil {
		log.Printf("[bot] failed to display note %s: %s", n.ID, err)
		return
	}
	if err := a.bot.Delete(m); err != nil {
		log.Printf("[bot] failed to delete original for note %s: %s", n.ID, err)
	}
}
