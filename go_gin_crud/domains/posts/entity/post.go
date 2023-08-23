package entity

import (
    "github.com/jinzhu/gorm"
)

type Post struct {
    gorm.Model
    Title   string
    Content string
    UserID  uint
}