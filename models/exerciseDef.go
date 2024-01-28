package models

type ExerciseDef struct {
	ExerciseDefId string `gorm:"primaryKey"`
	Name          string
}
