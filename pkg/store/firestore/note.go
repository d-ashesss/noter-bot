package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
)

const NoteCollection = "Notes"

type NoteStore struct {
	client *firestore.Client
}

func NewNoteStore(client *firestore.Client) *NoteStore {
	return &NoteStore{client: client}
}

func (s *NoteStore) Create(ctx context.Context, u interface{}) (string, error) {
	doc, _, err := s.client.Collection(NoteCollection).Add(ctx, u)
	if err != nil {
		return "", err
	}
	return doc.ID, nil
}

func (s *NoteStore) Get(ctx context.Context, id string, u interface{}) error {
	if id == "" {
		return errors.New("invalid note ID")
	}
	snap, err := s.client.Collection(NoteCollection).Doc(id).Get(ctx)
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
	if _, err := s.client.Collection(NoteCollection).Doc(id).Delete(ctx); err != nil {
		return err
	}
	return nil
}
