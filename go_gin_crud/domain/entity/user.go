package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `grom:"primaryKey"`
	Username string
	Email *string
	Age uint8
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}