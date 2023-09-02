package storage

import "github.com/sknutsen/Zdk/internal/data"

type ShoppingListsStorage struct {
	ZdkCtx *data.ZdkContext
}

func NewShoppingListsStorage(db *data.ZdkContext) *ShoppingListsStorage {
	return &ShoppingListsStorage{ZdkCtx: db}
}
