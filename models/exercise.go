package models

import "time"

type Exercise struct {
	ExerciseId    string `gorm:"primaryKey"`
	ExerciseDefId string
	WorkoutId     string
	UserId        string
	Name          string
	TotalDuration time.Duration
	Reps          uint
	Sets          uint
	Weight        float32
	DateCreated   time.Time
	DateChanged   time.Time
}
