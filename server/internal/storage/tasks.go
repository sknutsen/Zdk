package storage

import "github.com/sknutsen/Zdk/internal/data"

type TasksStorage struct {
	ZdkCtx *data.ZdkContext
}

func NewTasksStorage(db *data.ZdkContext) *TasksStorage {
	return &TasksStorage{ZdkCtx: db}
}
