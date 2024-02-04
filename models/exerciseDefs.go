package models

type Equipment struct {
	EquipmentId uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique"`
}

type ExerciseDef struct {
	ExerciseDefId uint   `gorm:"primaryKey"`
	Name          string `gorm:"unique"`
}

type UnitType struct {
	UnitTypeId uint   `gorm:"primaryKey"`
	Name       string `gorm:"unique"`
}
