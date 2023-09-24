package models

import (
	"time"

	"gorm.io/gorm"
)

type TaskCategory struct {
	gorm.Model
	TaskCategoryId string
	TaskId         string
	CategoryId     string
	DateCreated    time.Time
	DateChanged    time.Time
}

type DTOTaskCategoryListRequest struct {
	TaskId     string `json:"TaskId"`
	CategoryId string `json:"CategoryId"`
}

type DTOTaskCategoryListResponseData struct {
	TaskCategoryId string `json:"TaskCategoryId"`
	TaskId         string `json:"TaskId"`
	CategoryId     string `json:"CategoryId"`
}

type DTOTaskCategoryListResponse struct {
	List    []DTOTaskCategoryListResponseData `json:"List"`
	Message string                            `json:"Message"`
	Status  int                               `json:"Status"`
}

type DTOTaskCategoryNewRequest struct {
	TaskId     string `json:"TaskId"`
	CategoryId string `json:"CategoryId"`
}

type DTOTaskCategoryNewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOTaskCategoryUpdateRequest struct {
	TaskCategoryId string `json:"TaskCategoryId"`
	TaskId         string `json:"TaskId"`
	CategoryId     string `json:"CategoryId"`
}

type DTOTaskCategoryUpdateResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOTaskCategoryDeleteRequest struct {
	TaskCategoryId string `json:"TaskCategoryId"`
}

type DTOTaskCategoryDeleteResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}
