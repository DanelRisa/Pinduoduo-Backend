package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"pinduoduo-back/database"
	"pinduoduo-back/user-service/controllers"
	"pinduoduo-back/user-service/middleware"

	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
	
)

func main() {
	database.Connect()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:5173"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Authorization", "Content-Type"},
	}))
	
	r.Use(gin.RecoveryWithWriter(gin.DefaultErrorWriter))

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	
	r.Use(middleware.AuthMiddleware())
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	log.Println("User service running on http://localhost:8081")
	r.Run(":8081")
}
