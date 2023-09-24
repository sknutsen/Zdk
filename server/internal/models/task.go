package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskId      string
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
	Name string `json:"Name"`
}

type DTOTaskNewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOTaskUpdateRequest struct {
	TaskId string `json:"TaskId"`
	Name   string `json:"Name"`
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
