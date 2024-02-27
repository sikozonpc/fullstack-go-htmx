package handlers

import (
	"github.com/sikozonpc/fullstackgo/store"
)

type Handler struct {
	store *store.Storage
}

func New(store *store.Storage) *Handler {
	return &Handler{
		store: 		store,
	}
}