package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId      string
	DateCreated time.Time
	DateChanged time.Time
}

type UserProfile struct {
	UserId    string
	Name      string
	Nickname  string
	Picture   string
	Iat       string
	Exp       string
	UpdatedAt string
}

type DTOUserListRequest struct {
	UserId string `json:"UserId"`
}

type DTOUserListResponseData struct {
	UserId string `json:"UserId"`
}

type DTOUserListResponse struct {
	List    []DTOUserListResponseData `json:"List"`
	Message string                    `json:"Message"`
	Status  int                       `json:"Status"`
}

type DTOUserNewRequest struct {
	UserId string `json:"UserId"`
}

type DTOUserNewResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOUserUpdateRequest struct {
	UserId string `json:"UserId"`
}

type DTOUserUpdateResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

type DTOUserDeleteRequest struct {
	UserId string `json:"UserId"`
}

type DTOUserDeleteResponse struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}
