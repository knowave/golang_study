package repository

import (
	"context"

	"github.com/knowave/golang_study/tree/main/go_gin_crud/entity"
)

type UserRepositoryInterfacein interface {
	// Create user 생성
	Create(ctx context.Context, user *entity.User) (*entity.User, error)

	// GetAllByUser 모든 user 조회
	GetAllByUser(ctx context.Context, user *entity.User) (*entity.User, error)

	// GetUserById userId 조회
	GetUserById(ctx context.Context, userId uint) (*entity.User, error)

	// Update 특정 user update
	Update(ctx context.Context, userId uint) (*entity.User, error)

	// Delete 특정 User delete
	Delete(ctx context.Context, userId uint) (*entity.User, error)
}