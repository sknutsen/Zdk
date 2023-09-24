package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryId  string
	Name        string
	UserId      string
	DateCreated time.Time
	DateChanged time.Time
}

type DTOCategoryListRequest struct {
}

type DTOCategoryListResponseData struct {
	CategoryId string `json:"CategoryId"`
	Name       string `json:"Name"`
}

type DTOCategoryListResponse struct {
	List    []DTOCategoryListResponseData `json:"List"`
	Message string                        `json:"Message"`
	Status  int                           `json:"Status"`
}

type DTOCategoryNewRequest struct {
	Name string `json:"Name"`
}

type DTOCategoryNewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOCategoryUpdateRequest struct {
	CategoryId string `json:"CategoryId"`
	Name       string `json:"Name"`
}

type DTOCategoryUpdateResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOCategoryDeleteRequest struct {
	CategoryId string `json:"CategoryId"`
}

type DTOCategoryDeleteResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}
