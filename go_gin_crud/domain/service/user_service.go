package service

import (
	"context"

	"github.com/knowave/golang_study/tree/main/go_gin_crud/domain/entity"
	"github.com/knowave/golang_study/tree/main/go_gin_crud/domain/repository"
)

type UserServiceInterface interface {
	// GetAllByUser 유저 전체 조회
	GetAllByUser(ctx context.Context, user *entity.User) (*[]entity.User, error)

	// GetUserById 특정 유저 조회
	GetUserById(ctx context.Context, userId uint) (*entity.User, error)

	// CreateUser 유저 생성
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)

	// UpdateUser 특정 유저 수정
	UpdateUser(ctx context.Context, userId uint) (*entity.User, error)

	// Delete 특정 유저 삭제
	Delete(ctx context.Context, userId uint) (*entity.User, error)
}

type UserService struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserService(userRepo repository.UserRepositoryInterface) *UserService {
	return &UserService{userRepo}
}