package handlers

import (
	"github.com/sknutsen/Zdk/internal/storage"
)

type UserHandler struct {
	Storage *storage.ShoppingListItemsStorage
}

func NewUserHandler(storage *storage.ShoppingListItemsStorage) *UserHandler {
	return &UserHandler{Storage: storage}
}
