package store

import (
	"github.com/brenoandrade/estrategia/model"
)

// Result struct return of Stores
type Result struct {
	Data interface{}
	Err  *model.Error
}

// Channel type chan for return StoreResult
type Channel chan Result

// LayeredStore store layer
type LayeredStore struct {
	Store
}

// Do wrapper for encapsulate goroutine
func Do(f func(result *Result)) Channel {
	channel := make(Channel, 1)
	go func() {
		result := Result{}
		f(&result)
		channel <- result
		close(channel)
	}()
	return channel
}

// NewLayeredStore return instance db
func NewLayeredStore(db Store) Store {
	store := &LayeredStore{
		Store: db,
	}

	return store
}

// Close connection
func (s *LayeredStore) Close() {
	s.Store.Close()
}

// Store interface for define methods in App
type Store interface {
	Close()
	Crud() CrudStore
}

// CrudStore interface define methods
type CrudStore interface {
	Get(id string) Channel
}

// Crud return methods for CrudStore interface
func (s *LayeredStore) Crud() CrudStore {
	return s.Store.Crud()
}
