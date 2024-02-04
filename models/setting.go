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
