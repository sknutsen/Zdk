package storage

import "github.com/sknutsen/Zdk/internal/data"

type UsersStorage struct {
	ZdkCtx *data.ZdkContext
}

func NewUsersStorage(db *data.ZdkContext) *UsersStorage {
	return &UsersStorage{ZdkCtx: db}
}
