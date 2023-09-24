package storage

import "github.com/sknutsen/Zdk/internal/data"

type CategoriesStorage struct {
	ZdkCtx *data.ZdkContext
}

func NewCategoriesStorage(db *data.ZdkContext) *CategoriesStorage {
	return &CategoriesStorage{ZdkCtx: db}
}
