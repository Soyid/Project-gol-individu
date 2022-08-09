package Model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          int `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Description string
}
