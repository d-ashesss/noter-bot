package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/d-ashesss/noter-bot/pkg/model"
)

const NoteCollectionName = "Notes"

type NoteStore struct {
	client *firestore.Client
}

func NewNoteStore(client *firestore.Client) *NoteStore {
	return &NoteStore{client: client}
}

func (s *NoteStore) Create(ctx context.Context, u interface{}) (string, error) {
	doc, _, err := s.client.Collection(NoteCollectionName).Add(ctx, u)
	if err != nil {
		return "", err
	}
	return doc.ID, nil
}

func (s *NoteStore) Get(ctx context.Context, id string, u interface{}) error {
	if id == "" {
		return errors.New("invalid note ID")
	}
	snap, err := s.client.Collection(NoteCollectionName).Doc(id).Get(ctx)
	if err != nil {
		return err
	}
	if err := snap.DataTo(u); err != nil {
		return err
	}
	return nil
}

func (s *NoteStore) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("invalid note ID")
	}
	if _, err := s.client.Collection(NoteCollectionName).Doc(id).Delete(ctx); err != nil {
		return err
	}
	return nil
}

func (s *NoteStore) FindByUser(userID int64) model.NoteCollection {
	return &NoteCollection{
		query: s.client.Collection(NoteCollectionName).Where(model.NoteFieldUserID, "==", userID),
	}
}

type NoteCollection struct {
	query firestore.Query
}

func (c *NoteCollection) All(ctx context.Context) <-chan *model.Note {
	notes := make(chan *model.Note, 1)
	iter := c.query.Documents(ctx)
	go func() {
		defer close(notes)

		for {
			snap, err := iter.Next()
			if err != nil {
				break
			}
			var n model.Note
			if err := snap.DataTo(&n); err != nil {
				continue
			}
			n.ID = snap.Ref.ID
			notes <- &n
		}
	}()
	return notes
}
