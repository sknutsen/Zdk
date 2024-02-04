package models

import "time"

type Workout struct {
	WorkoutId   string `gorm:"primaryKey"`
	Name        string
	Description string
	Date        string
	UserId      string
	DateCreated time.Time
	DateChanged time.Time
}

type DTOWorkoutNewRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type DTOWorkoutUpdateRequest struct {
	WorkoutId   string     `json:"workoutId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Date        string     `json:"date"`
	Exercises   []Exercise `json:"exercises"`
}
