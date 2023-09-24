package storage

import "github.com/sknutsen/Zdk/internal/data"

type TaskCategoriesStorage struct {
	ZdkCtx *data.ZdkContext
}

func NewTaskCategoriesStorage(db *data.ZdkContext) *TaskCategoriesStorage {
	return &TaskCategoriesStorage{ZdkCtx: db}
}
