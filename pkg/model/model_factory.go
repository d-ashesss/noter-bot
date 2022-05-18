package model

import (
	"cloud.google.com/go/firestore"
	store "github.com/d-ashesss/noter-bot/pkg/store/firestore"
)

type Factory interface {
	NewNoteModel() *NoteModel
}

func NewFirestoreFactory(storeClient *firestore.Client) Factory {
	return &FirestoreFactory{storeClient: storeClient}
}

type FirestoreFactory struct {
	storeClient *firestore.Client
}

func (f *FirestoreFactory) NewNoteModel() *NoteModel {
	s := store.NewNoteStore(f.storeClient)
	return NewNoteModel(s)
}
