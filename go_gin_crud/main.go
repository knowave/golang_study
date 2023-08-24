package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/knowave/golang_study/tree/main/go_gin_crud/go_gin_crud/config"
	"github.com/knowave/golang_study/tree/main/go_gin_crud/go_gin_crud/domains/posts"
	"github.com/knowave/golang_study/tree/main/go_gin_crud/go_gin_crud/domains/users"
)

func main() {
	loadErr := godotenv.Load(".env")

	if loadErr != nil {
		panic("Error Loading .env File")
	}

	port := os.Getenv("PORT")


	config.Init()

	postRepo := posts.NewPostRepository(config.DB)
	userRepo := users.NewUserRepository(config.DB)

	postService := posts.NewPostService(postRepo)
	userService := users.NewUserService(userRepo)

	_ = posts.NewPostController(postService)
	_ = users.NewUserController(userService)

	r := gin.Default()

	
	r.Run(port)
}