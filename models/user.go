package models

import (
	"time"
)

type User struct {
	UserId      string
	DateCreated time.Time
	DateChanged time.Time
}

type UserProfile struct {
	UserId    string
	Name      string
	Nickname  string
	Picture   string
	Iat       float64
	Exp       float64
	UpdatedAt string
	Admin     bool
}

func GetUserProfile(profile map[string]interface{}) UserProfile {
	return UserProfile{
		UserId:    profile["sub"].(string),
		Name:      profile["name"].(string),
		Picture:   profile["picture"].(string),
		Nickname:  profile["nickname"].(string),
		Iat:       profile["iat"].(float64),
		Exp:       profile["exp"].(float64),
		UpdatedAt: profile["updated_at"].(string),
	}
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
