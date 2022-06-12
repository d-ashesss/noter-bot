package model

import (
	"context"
	"fmt"
	"time"
)

const (
	NoteFieldUserID = "UserID"
	NoteFieldDate   = "Date"
)

// Note represents a category entity.
type Note struct {
	ID     string `firestore:"-"`
	UserID int64
	Text   string
	Date   time.Time
}

// NewNote initializes new Note.
func NewNote(userID int64, text string) *Note {
	return &Note{
		UserID: userID,
		Text:   text,
		Date:   time.Now(),
	}
}

// NoteStore is an interface wrapper for a DB engine.
type NoteStore interface {
	Create(ctx context.Context, u interface{}) (string, error)
	Get(ctx context.Context, id string, u interface{}) error
	Delete(ctx context.Context, id string) error
	FindByUser(userID int64) NoteCollection
}

// NoteModel data model for Note.
type NoteModel struct {
	db NoteStore
}

// NewNoteModel initializes NoteModel.
func NewNoteModel(db NoteStore) *NoteModel {
	return &NoteModel{db: db}
}

func (m *NoteModel) Create(ctx context.Context, n *Note) (err error) {
	if len(n.ID) > 0 {
		return fmt.Errorf("create categore: provided categore is not new")
	}
	n.ID, err = m.db.Create(ctx, n)
	return
}

func (m *NoteModel) Get(ctx context.Context, id string) (*Note, error) {
	var n Note
	if err := m.db.Get(ctx, id, &n); err != nil {
		return nil, err
	}
	n.ID = id
	return &n, nil
}

func (m *NoteModel) Delete(ctx context.Context, n *Note) error {
	return m.db.Delete(ctx, n.ID)
}

func (m *NoteModel) FindByUser(userID int64) NoteCollection {
	c := m.db.FindByUser(userID)
	return c
}

type NoteCollection interface {
	All(ctx context.Context) <-chan *Note
	First(ctx context.Context) *Note
	Last(ctx context.Context) *Note
}
