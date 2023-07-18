package persistence

import (
	"context"

	"github.com/knowave/golang_study/tree/main/go_gin_crud/domain/entity"
	"github.com/knowave/golang_study/tree/main/go_gin_crud/domain/repository"
	"gorm.io/gorm"
)

var _ repository.UserRepositoryInterface = &userRepo{}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

	// Create user 생성
	func (u *userRepo) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
		return user, nil
	}

	// GetAllByUser 모든 user 조회
	func (u *userRepo) GetAllByUser(ctx context.Context, user *entity.User) (*[]entity.User, error) {
		var users *[]entity.User
		return users, nil
	}

	// GetUserById userId 조회
	func (u *userRepo) GetUserById(ctx context.Context, userId uint) (*entity.User, error) {
		var user *entity.User
		return user, nil
	}

	// Update 특정 user update
	func (u *userRepo) Update(ctx context.Context, userId uint) (*entity.User, error) {
		var user *entity.User
		return user, nil
	}

	// Delete 특정 User delete
	func (u *userRepo) Delete(ctx context.Context, userId uint) (*entity.User, error) {
		var user *entity.User
		return user, nil
	}