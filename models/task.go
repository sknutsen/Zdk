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

type DTOTaskListRequest struct {
}

type DTOTaskListResponseData struct {
	TaskId string `json:"TaskId"`
	Name   string `json:"Name"`
}

type DTOTaskListResponse struct {
	List    []DTOTaskListResponseData `json:"List"`
	Message string                    `json:"Message"`
	Status  int                       `json:"Status"`
}

type DTOTaskNewRequest struct {
	Name       string   `json:"name"`
	Categories []string `json:"categories"`
}

type DTOTaskNewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOTaskUpdateRequest struct {
	TaskId     string   `json:"taskId"`
	Name       string   `json:"name"`
	Categories []string `json:"categories"`
}

type DTOTaskUpdateResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOTaskDeleteRequest struct {
	TaskId string `json:"TaskId"`
}

type DTOTaskDeleteResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}
