package models

import (
	"time"
)

type Task struct {
	TaskId      string `gorm:"primaryKey"`
	Name        string
	UserId      string
	DateCreated time.Time
	DateChanged time.Time
}

type DTOTaskNewRequest struct {
	Name       string   `json:"name"`
	Categories []string `json:"categories"`
}

type DTOTaskUpdateRequest struct {
	TaskId     string   `json:"taskId"`
	Name       string   `json:"name"`
	Categories []string `json:"categories"`
}
