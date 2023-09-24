package storage

import "github.com/sknutsen/Zdk/internal/data"

type SettingsStorage struct {
	ZdkCtx *data.ZdkContext
}

func NewSettingsStorage(db *data.ZdkContext) *SettingsStorage {
	return &SettingsStorage{ZdkCtx: db}
}
