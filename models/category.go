package models

import (
	"time"
)

type Category struct {
	CategoryId  string `gorm:"primaryKey"`
	Name        string
	UserId      string
	DateCreated time.Time
	DateChanged time.Time
}

type DTOCategoryListRequest struct {
}

type DTOCategoryListResponseData struct {
	CategoryId string `json:"categoryId"`
	Name       string `json:"Name"`
}

type DTOCategoryListResponse struct {
	List    []DTOCategoryListResponseData `json:"List"`
	Message string                        `json:"Message"`
	Status  int                           `json:"Status"`
}

type DTOCategoryNewRequest struct {
	Name  string   `json:"name"`
	Tasks []string `json:"tasks"`
}

type DTOCategoryNewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOCategoryUpdateRequest struct {
	CategoryId string   `json:"categoryId"`
	Name       string   `json:"name"`
	Tasks      []string `json:"tasks"`
}

type DTOCategoryUpdateResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOCategoryDeleteRequest struct {
	CategoryId string `json:"categoryId"`
}

type DTOCategoryDeleteResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}
