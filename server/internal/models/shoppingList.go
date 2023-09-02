package models

import "gorm.io/gorm"

type ShoppingList struct {
	gorm.Model
	ShoppingListId string
	Name           string
}

type ShoppingListDTOListRequest struct {
}

type ShoppingListDTOListResponseData struct {
	ShoppingListId string `json:"ShoppingListId"`
	Name           string `json:"Name"`
}

type ShoppingListDTOListResponse struct {
	List    []ShoppingListDTOListResponseData `json:"List"`
	Message string                            `json:"Message"`
	Status  int                               `json:"Status"`
}

type ShoppingListDTONewRequest struct {
	Name string `json:"Name"`
}

type ShoppingListDTONewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type ShoppingListDTOUpdateRequest struct {
	ShoppingListId string `json:"ShoppingListId"`
	Name           string `json:"Name"`
}

type ShoppingListDTOUpdateResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type ShoppingListDTODeleteRequest struct {
	ShoppingListId string `json:"ShoppingListId"`
}

type ShoppingListDTODeleteResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}
