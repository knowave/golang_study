package entity

import (
	"github.com/jinzhu/gorm"
	"github.com/knowave/golang_study/tree/main/go_gin_crud/go_gin_crud/domains/posts/entity"
)

type User struct {
    gorm.Model
    Username string
    Email    string
    Password string
    Posts    []entity.Post
}