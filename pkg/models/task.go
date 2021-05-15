package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	TaskText  string
	TaskCode  string
}
