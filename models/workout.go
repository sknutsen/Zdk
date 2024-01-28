package models

import "time"

type Workout struct {
	WorkoutId   string `gorm:"primaryKey"`
	Name        string
	Date        time.Time
	UserId      string
	DateCreated time.Time
	DateChanged time.Time
}
