// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/d-ashesss/noter-bot/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// NoteCollection is an autogenerated mock type for the NoteCollection type
type NoteCollection struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx
func (_m *NoteCollection) All(ctx context.Context) <-chan *model.Note {
	ret := _m.Called(ctx)

	var r0 <-chan *model.Note
	if rf, ok := ret.Get(0).(func(context.Context) <-chan *model.Note); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *model.Note)
		}
	}

	return r0
}