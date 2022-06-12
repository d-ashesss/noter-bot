package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/d-ashesss/noter-bot/pkg/model"
	"google.golang.org/api/iterator"
	"log"
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
		query: s.client.Collection(NoteCollectionName).
			Where(model.NoteFieldUserID, "==", userID).
			OrderBy(model.NoteFieldDate, firestore.Asc),
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
		defer iter.Stop()

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

func (c *NoteCollection) First(ctx context.Context) *model.Note {
	iter := c.query.Limit(1).Documents(ctx)
	defer iter.Stop()
	snap, err := iter.Next()
	if err != nil {
		if err != iterator.Done {
			log.Printf("[firestore] failed to get first note: %s", err)
		}
		return nil
	}
	var n model.Note
	if err := snap.DataTo(&n); err != nil {
		return nil
	}
	n.ID = snap.Ref.ID
	return &n
}

func (c *NoteCollection) Last(ctx context.Context) *model.Note {
	iter := c.query.LimitToLast(1).Documents(ctx)
	defer iter.Stop()
	snaps, err := iter.GetAll()
	if err != nil {
		log.Printf("[firestore] failed to get last note: %s", err)
		return nil
	}
	var n model.Note
	if err := snaps[0].DataTo(&n); err != nil {
		return nil
	}
	n.ID = snaps[0].Ref.ID
	return &n
}
