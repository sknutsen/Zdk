package models

import (
	"time"

	"gorm.io/gorm"
)

type ScheduledTask struct {
	gorm.Model
	ScheduledTaskId string
	TaskId          string
	IsComplete      bool
	Date            time.Time
	DateCreated     time.Time
	DateChanged     time.Time
}

type DTOScheduledTaskListRequest struct {
	TaskId     string    `json:"TaskId"`
	IsComplete bool      `json:"IsComplete"`
	Date       time.Time `json:"Date"`
}

type DTOScheduledTaskListResponseData struct {
	ScheduledTaskId string `json:"ScheduledTaskId"`
}

type DTOScheduledTaskListResponse struct {
	List    []DTOScheduledTaskListResponseData `json:"List"`
	Message string                             `json:"Message"`
	Status  int                                `json:"Status"`
}

type DTOScheduledTaskNewRequest struct {
	TaskId     string    `json:"TaskId"`
	Name       string    `json:"Name"`
	IsComplete bool      `json:"IsComplete"`
	Date       time.Time `json:"Date"`
}

type DTOScheduledTaskNewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOScheduledTaskUpdateRequest struct {
	ScheduledTaskId string    `json:"ScheduledTaskId"`
	Name            string    `json:"Name"`
	IsComplete      bool      `json:"IsComplete"`
	Date            time.Time `json:"Date"`
}

type DTOScheduledTaskUpdateResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOScheduledTaskDeleteRequest struct {
	ScheduledTaskId string `json:"ScheduledTaskId"`
}

type DTOScheduledTaskDeleteResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}
