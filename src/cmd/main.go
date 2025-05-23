package main

import (
	"net/http"
	"user-management/src/config"
	"user-management/src/controller"
	"user-management/src/repository"
	"user-management/src/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	healthController := controller.NewHealthController()

	// Configuração do Gin e Rotas
	r := gin.Default()
	r.POST("/user", userController.CreateUser)
	r.GET("/users/:id", userController.GetUserById)
	r.GET("/users", userController.GetAllUsers)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	//  Readines and liveness Routers

	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	r.GET("/readyz", healthController.Readiness)
	r.GET("/readyz", healthController.Liveness)
	r.Run(":8080")
}
