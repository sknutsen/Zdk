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

type DTOCategoryNewRequest struct {
	Name  string   `json:"name"`
	Tasks []string `json:"tasks"`
}

type DTOCategoryUpdateRequest struct {
	CategoryId string   `json:"categoryId"`
	Name       string   `json:"name"`
	Tasks      []string `json:"tasks"`
}
