package storage

import "github.com/sknutsen/Zdk/internal/data"

type ScheduledTasksStorage struct {
	ZdkCtx *data.ZdkContext
}

func NewScheduledTasksStorage(db *data.ZdkContext) *ScheduledTasksStorage {
	return &ScheduledTasksStorage{ZdkCtx: db}
}
