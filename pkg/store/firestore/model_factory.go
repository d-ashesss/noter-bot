package firestore

import (
	"cloud.google.com/go/firestore"
	"github.com/d-ashesss/noter-bot/pkg/model"
)

func NewFactory(storeClient *firestore.Client) model.Factory {
	return &Factory{storeClient: storeClient}
}

type Factory struct {
	storeClient *firestore.Client
}

func (f *Factory) NewNoteModel() *model.NoteModel {
	s := NewNoteStore(f.storeClient)
	return model.NewNoteModel(s)
}
