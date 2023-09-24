package models

import (
	"time"

	"gorm.io/gorm"
)

type ShoppingListItem struct {
	gorm.Model
	ShoppingListItemId string
	ShoppingListId     string
	Name               string
	DateCreated        time.Time
	DateChanged        time.Time
}

type DTOShoppingListItemListRequest struct {
	ShoppingListId string `json:"ShoppingListId"`
}

type DTOShoppingListItemListResponseData struct {
	ShoppingListItemId string `json:"ShoppingListItemId"`
	ShoppingListId     string `json:"ShoppingListId"`
	Name               string `json:"Name"`
}

type DTOShoppingListItemListResponse struct {
	List    []DTOShoppingListItemListResponseData `json:"List"`
	Message string                                `json:"Message"`
	Status  int                                   `json:"Status"`
}

type DTOShoppingListItemNewRequest struct {
	ShoppingListId string `json:"ShoppingListId"`
	Name           string `json:"Name"`
}

type DTOShoppingListItemNewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOShoppingListItemUpdateRequest struct {
	ShoppingListItemId string `json:"ShoppingListItemId"`
	Name               string `json:"Name"`
}

type DTOShoppingListItemUpdateResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOShoppingListItemDeleteRequest struct {
	ShoppingListItemId string `json:"ShoppingListItemId"`
}

type DTOShoppingListItemDeleteResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}
