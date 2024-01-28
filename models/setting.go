package models

import (
	"time"
)

type Setting struct {
	SettingId   string `gorm:"primaryKey"`
	Name        string
	Type        string
	DateCreated time.Time
	DateChanged time.Time
}

type DTOSettingListRequest struct {
}

type DTOSettingListResponseData struct {
	SettingId string `json:"SettingId"`
	Name      string `json:"Name"`
	Type      string `json:"Type"`
}

type DTOSettingListResponse struct {
	List    []DTOSettingListResponseData `json:"List"`
	Message string                       `json:"Message"`
	Status  int                          `json:"Status"`
}

type DTOSettingNewRequest struct {
	Name string `json:"Name"`
	Type string `json:"Type"`
}

type DTOSettingNewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOSettingUpdateRequest struct {
	SettingId string `json:"SettingId"`
	Name      string `json:"Name"`
	Type      string `json:"Type"`
}

type DTOSettingUpdateResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOSettingDeleteRequest struct {
	SettingId string `json:"SettingId"`
}

type DTOSettingDeleteResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}
