package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/knowave/golang_study/tree/main/go_gin_crud/domain/service"
	"github.com/knowave/golang_study/tree/main/go_gin_crud/infra/persistence"
)

var (
    dbHost     string
	dbPassword string
	dbUsername string
	dbName     string
	dbPort     string
)

func init() {
    loadErr := godotenv.Load()

    if loadErr != nil {
        fmt.Printf("end error %+v", loadErr)
    }

    dbHost = os.Getenv("DB_HOST")
    dbPassword = os.Getenv("DB_PASSWORD")
    dbUsername = os.Getenv("DB_USERNAME")
    dbName = os.Getenv("DB_NAME")
    dbPort = os.Getenv("DB_PORT")
}

func main() {
    repositories, err := persistence.NewRepositories(dbHost, dbPort, dbUsername, dbName, dbPassword)
	if err != nil {
		fmt.Errorf("[Error] persistence.NewRepositories  %v", err)
		return
	}

    userService := service.NewUserService(repositories)

    r := gin.Default()
    r.GET("", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

    r.Run("localhost:3000")
}