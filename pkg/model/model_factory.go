package model

type Factory interface {
	NewNoteModel() *NoteModel
}
