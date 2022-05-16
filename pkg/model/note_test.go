package model

import (
	"context"
	"errors"
	"github.com/d-ashesss/noter-bot/mocks"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	testNoteError = errors.New("test error")
)

const (
	testNoteID = "test-ID"
	testUserID = 666
)

func TestNoteModel_Create(t *testing.T) {
	t.Run("user exists", func(t *testing.T) {
		db := &mocks.NoteStore{}

		m := NewNoteModel(db)
		n := NewNote(testUserID, "")
		n.ID = "new-ID"

		if err := m.Create(context.Background(), n); err == nil {
			t.Errorf("Create() expected error")
		}
	})

	t.Run("create error", func(t *testing.T) {
		db := &mocks.NoteStore{}
		db.On("Create", mock.Anything, mock.Anything).Return("", testNoteError)

		m := NewNoteModel(db)
		n := NewNote(testUserID, "")

		if err := m.Create(context.Background(), n); err != testNoteError {
			t.Errorf("Create() got error %q, want %q", err, testNoteError)
		}
	})

	t.Run("user created", func(t *testing.T) {
		db := &mocks.NoteStore{}
		db.On("Create", mock.Anything, mock.Anything).Return(func(ctx context.Context, o interface{}) string {
			return testNoteID
		}, nil)

		m := NewNoteModel(db)
		n := NewNote(testUserID, "")

		if err := m.Create(context.Background(), n); err != nil {
			t.Errorf("Create() unexpected error = %v", err)
		}
		if n.ID != testNoteID {
			t.Errorf("Create() got user ID %q, want %q", n.ID, testNoteID)
		}
	})
}

func TestNoteModel_Get(t *testing.T) {
	t.Run("user not found", func(t *testing.T) {
		db := &mocks.NoteStore{}
		db.On("Get", mock.Anything, mock.Anything, mock.Anything).Return(testNoteError)
		m := NewNoteModel(db)

		n, err := m.Get(context.Background(), testNoteID)
		if err != testNoteError {
			t.Errorf("Get() got error %q, want %q", err, testNoteError)
		}
		if n != nil {
			t.Errorf("Get() returned unexpected user")
		}
	})

	t.Run("user found", func(t *testing.T) {
		db := &mocks.NoteStore{}
		db.On("Get", mock.Anything, mock.MatchedBy(func(id string) bool {
			return id == testNoteID
		}), mock.Anything).Return(nil)
		m := NewNoteModel(db)

		n, err := m.Get(context.Background(), testNoteID)
		if err != nil {
			t.Errorf("Get() unexpected error = %v", err)
		}
		if n.ID != testNoteID {
			t.Errorf("Get() got user with ID %q, want %q", n.ID, testNoteID)
		}
	})
}

func TestNoteModel_Delete(t *testing.T) {
	t.Run("category deleted", func(t *testing.T) {
		db := &mocks.NoteStore{}
		db.On("Delete", mock.Anything, mock.MatchedBy(func(id string) bool {
			return id == testNoteID
		})).Return(nil)
		m := NewNoteModel(db)
		n := NewNote(testUserID, "")
		n.ID = testNoteID
		if err := m.Delete(context.Background(), n); err != nil {
			t.Errorf("Delete() unexpected error = %v", err)
		}
	})
}
