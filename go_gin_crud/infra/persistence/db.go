package persistence

import (
	"fmt"

	"github.com/knowave/golang_study/tree/main/go_gin_crud/domain/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repositories struct {
	repository.UserRepositoryInterface
}

func NewRepositories(host, port, user, dbname, password string) (*Repositories, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed database connection: %w", err)
	}

	return &Repositories{
		NewUserRepository(db),
	}, nil
}