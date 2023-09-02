package models

import "gorm.io/gorm"

type ShoppingListItem struct {
	gorm.Model
	ShoppingListItemId string
	ShoppingListId     string
	Name               string
}

type ShoppingListItemDTOListRequest struct {
	ShoppingListId string `json:"ShoppingListId"`
}

type ShoppingListItemDTOListResponseData struct {
	ShoppingListItemId string `json:"ShoppingListItemId"`
	ShoppingListId     string `json:"ShoppingListId"`
	Name               string `json:"Name"`
}

type ShoppingListItemDTOListResponse struct {
	List    []ShoppingListDTOListResponseData `json:"List"`
	Message string                            `json:"Message"`
	Status  int                               `json:"Status"`
}

type ShoppingListItemDTONewRequest struct {
	ShoppingListId string `json:"ShoppingListId"`
	Name           string `json:"Name"`
}

type ShoppingListItemDTONewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type ShoppingListItemDTOUpdateRequest struct {
	ShoppingListItemId string `json:"ShoppingListItemId"`
	Name               string `json:"Name"`
}

type ShoppingListItemDTOUpdateResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type ShoppingListItemDTODeleteRequest struct {
	ShoppingListItemId string `json:"ShoppingListItemId"`
}

type ShoppingListItemDTODeleteResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}
