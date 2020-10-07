package main

import (
	"user-api/config"
	"user-api/controllers"
	"user-api/handler"
	"user-api/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	db := config.Connect()

	repo := repositories.NewUserRepository(db)
	entity := controllers.NewUserEntity(repo)

	api := r.Group("/api")
	handler.UserHandlerFunc(api, entity)
	r.Run()
}
