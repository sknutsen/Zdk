package storage

import "github.com/sknutsen/Zdk/internal/data"

type ShoppingListItemsStorage struct {
	ZdkCtx *data.ZdkContext
}

func NewShoppingListItemsStorage(db *data.ZdkContext) *ShoppingListItemsStorage {
	return &ShoppingListItemsStorage{ZdkCtx: db}
}
