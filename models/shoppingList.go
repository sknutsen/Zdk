package models

import (
	"time"
)

type ShoppingList struct {
	ShoppingListId string `gorm:"primaryKey"`
	Name           string
	DateCreated    time.Time
	DateChanged    time.Time
}

type DTOShoppingListListRequest struct {
}

type DTOShoppingListListResponseData struct {
	ShoppingListId string `json:"ShoppingListId"`
	Name           string `json:"Name"`
}

type DTOShoppingListListResponse struct {
	List    []DTOShoppingListListResponseData `json:"List"`
	Message string                            `json:"Message"`
	Status  int                               `json:"Status"`
}

type DTOShoppingListNewRequest struct {
	Name string `json:"Name"`
}

type DTOShoppingListNewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOShoppingListUpdateRequest struct {
	ShoppingListId string `json:"ShoppingListId"`
	Name           string `json:"Name"`
}

type DTOShoppingListUpdateResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOShoppingListDeleteRequest struct {
	ShoppingListId string `json:"ShoppingListId"`
}

type DTOShoppingListDeleteResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}
