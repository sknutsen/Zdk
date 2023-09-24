package storage

import "github.com/sknutsen/Zdk/internal/data"

type UserSettingsStorage struct {
	ZdkCtx *data.ZdkContext
}

func NewUserSettingsStorage(db *data.ZdkContext) *UserSettingsStorage {
	return &UserSettingsStorage{ZdkCtx: db}
}
