package models

import (
	"time"
)

type UserSetting struct {
	UserSettingId string `gorm:"primaryKey"`
	UserId        string
	SettingId     string
	Value         string
	DateCreated   time.Time
	DateChanged   time.Time
}

type DTOUserSettingListRequest struct {
	UserSettingId string `json:"UserSettingId"`
}

type DTOUserSettingListResponseData struct {
	UserSettingId string `json:"UserSettingId"`
	UserId        string `json:"UserId"`
	SettingId     string `json:"SettingId"`
	Value         string `json:"Value"`
}

type DTOUserSettingListResponse struct {
	List    []DTOUserSettingListResponseData `json:"List"`
	Message string                           `json:"Message"`
	Status  int                              `json:"Status"`
}

type DTOUserSettingNewRequest struct {
	SettingId string `json:"SettingId"`
	Value     string `json:"Value"`
}

type DTOUserSettingNewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOUserSettingUpdateRequest struct {
	UserSettingId string `json:"UserSettingId"`
	SettingId     string `json:"SettingId"`
	Value         string `json:"Value"`
}

type DTOUserSettingUpdateResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOUserSettingDeleteRequest struct {
	UserSettingId string `json:"UserSettingId"`
}

type DTOUserSettingDeleteResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}
