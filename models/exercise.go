package models

import "time"

type Exercise struct {
	ExerciseId    string `gorm:"primaryKey"`
	ExerciseDefId uint
	EquipmentId   uint
	WorkoutId     string
	UserId        string
	Name          string
	Hours         uint
	Minutes       uint
	Seconds       uint
	Units         uint
	UnitTypeId    uint
	Sets          uint
	Weight        float32
	DateCreated   time.Time
	DateChanged   time.Time
}

type DTOExerciseNewRequest struct {
	ExerciseDefId uint    `json:"exerciseDefId"`
	EquipmentId   uint    `json:"equipmentId"`
	WorkoutId     string  `json:"workoutId"`
	Name          string  `json:"name"`
	Hours         uint    `json:"hours"`
	Minutes       uint    `json:"minutes"`
	Seconds       uint    `json:"seconds"`
	Units         uint    `json:"units"`
	UnitTypeId    uint    `json:"unitTypeId"`
	Sets          uint    `json:"sets"`
	Weight        float32 `json:"weight"`
}

type DTOExerciseUpdateRequest struct {
	ExerciseId    string  `json:"exerciseId"`
	ExerciseDefId uint    `json:"exerciseDefId"`
	EquipmentId   uint    `json:"equipmentId"`
	WorkoutId     string  `json:"workoutId"`
	Name          string  `json:"name"`
	Hours         uint    `json:"hours"`
	Minutes       uint    `json:"minutes"`
	Seconds       uint    `json:"seconds"`
	Units         uint    `json:"units"`
	UnitTypeId    uint    `json:"unitTypeId"`
	Sets          uint    `json:"sets"`
	Weight        float32 `json:"weight"`
}
