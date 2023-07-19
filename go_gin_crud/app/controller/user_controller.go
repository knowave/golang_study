package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/knowave/golang_study/tree/main/go_gin_crud/domain/repository"
	"github.com/knowave/golang_study/tree/main/go_gin_crud/domain/service"
)

type UserController struct {
	userService service.UserService
	userRepo repository.UserRepositoryInterface
	gin *gin.Context
}

func NewUserController(userService service.UserService, userRepo repository.UserRepositoryInterface, gin *gin.Context) *UserController {
	return &UserController{userService: userService, userRepo: userRepo, gin: gin}
}

type UserControllerInterface interface {
	// SignUp 회원가입
	SignUp(gin *gin.Context)

	// GetAllByUser 유저 전체 조회
	GetAllByUser(gin *gin.Context)

	// GetUserById 특정 유저 조회
	GetUserById(gin *gin.Context)

	// UpdateUser 특정 유저 정보 수정
	UpdateUser(gin *gin.Context)

	// DeleteUser 특정 유저 삭제
	DeleteUser(gin *gin.Context)
}