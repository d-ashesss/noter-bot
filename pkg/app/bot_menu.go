package app

import "gopkg.in/tucnak/telebot.v2"

const (
	BotMenuNoteOptionsBtnDeleteLabel = "Delete"
	BotMenuNoteOptionsBtnDeleteID    = "btnMenuNoteOptionsDelete"
)

// BotMenuNoteOptions represents Note options menu.
type BotMenuNoteOptions struct {
	Menu *telebot.ReplyMarkup

	BtnDelete telebot.Btn
}

// NewBotMenuNoteOptions initializes new BotMenuNoteOptions.
func NewBotMenuNoteOptions(noteID string) *BotMenuNoteOptions {
	m := &BotMenuNoteOptions{
		Menu: &telebot.ReplyMarkup{},
	}
	m.BtnDelete = m.Menu.Data(BotMenuNoteOptionsBtnDeleteLabel, BotMenuNoteOptionsBtnDeleteID, noteID)
	m.Menu.Inline(
		m.Menu.Row(m.BtnDelete),
	)
	return m
}

const (
	BotMenuMyNotesBtnShowAllLabel   = "Show all"
	BotMenuMyNotesBtnShowAllID      = "btnMenuMyNotesShowAll"
	BotMenuMyNotesBtnShowFirstLabel = "Show first"
	BotMenuMyNotesBtnShowFirstID    = "btnMenuMyNotesShowFirst"
	BotMenuMyNotesBtnShowLastLabel  = "Show last"
	BotMenuMyNotesBtnShowLastID     = "btnMenuMyNotesShowLast"
)

// BotMenuMyNotes represents My notes menu.
type BotMenuMyNotes struct {
	Menu *telebot.ReplyMarkup

	BtnShowAll   telebot.Btn
	BtnShowFirst telebot.Btn
	BtnShowLast  telebot.Btn
}

// NewBotMenuMyNotes initializes new BotMenuMyNotes.
func NewBotMenuMyNotes() *BotMenuMyNotes {
	m := &BotMenuMyNotes{
		Menu: &telebot.ReplyMarkup{},
	}
	m.BtnShowAll = m.Menu.Data(BotMenuMyNotesBtnShowAllLabel, BotMenuMyNotesBtnShowAllID)
	m.BtnShowFirst = m.Menu.Data(BotMenuMyNotesBtnShowFirstLabel, BotMenuMyNotesBtnShowFirstID)
	m.BtnShowLast = m.Menu.Data(BotMenuMyNotesBtnShowLastLabel, BotMenuMyNotesBtnShowLastID)
	m.Menu.Inline(
		m.Menu.Row(m.BtnShowAll),
		m.Menu.Row(m.BtnShowFirst, m.BtnShowLast),
	)
	return m
}
