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
}