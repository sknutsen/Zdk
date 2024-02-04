package models

import (
	"time"
)

type TaskCategory struct {
	TaskCategoryId string `gorm:"primaryKey"`
	TaskId         string
	CategoryId     string
	DateCreated    time.Time
	DateChanged    time.Time
}
