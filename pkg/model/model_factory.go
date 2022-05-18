package model

import "cloud.google.com/go/firestore"

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
	//TODO implement me
	panic("implement me")
}
