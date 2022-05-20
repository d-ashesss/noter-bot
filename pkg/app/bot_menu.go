package app

import "gopkg.in/tucnak/telebot.v2"

const (
	botMenuNoteOptionsBtnDeleteLabel = "Delete"
	botMenuNoteOptionsBtnDeleteID    = "btnMenuNoteOptionsDelete"
)

// BotMenuNoteOptions represents Note options menu.
type BotMenuNoteOptions struct {
	Menu *telebot.ReplyMarkup

	BtnDelete *telebot.Btn
}

// NewBotMenuNoteOptions initializes new BotMenuNoteOptions.
func NewBotMenuNoteOptions(noteID string) *BotMenuNoteOptions {
	m := &BotMenuNoteOptions{
		Menu: &telebot.ReplyMarkup{},
	}
	btnDelete := m.Menu.Data(botMenuNoteOptionsBtnDeleteLabel, botMenuNoteOptionsBtnDeleteID, noteID)
	m.BtnDelete = &btnDelete
	m.Menu.Inline(m.Menu.Row(btnDelete))
	return m
}
