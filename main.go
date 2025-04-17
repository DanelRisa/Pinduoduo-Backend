package main

import (
	"log"
	"pinduoduo-back/controllers"
	"pinduoduo-back/database"
	"pinduoduo-back/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.Migrate()

	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/products", controllers.GetProducts)
		auth.GET("/products/:id", controllers.GetProduct)
		auth.POST("/products", controllers.CreateProduct)
		auth.PUT("/products/:id", controllers.UpdateProduct)
		auth.DELETE("/products/:id", controllers.DeleteProduct)

		auth.POST("/groupbuys", controllers.CreateGroupBuy)
		auth.GET("/groupbuys", controllers.GetGroupBuys)
		auth.GET("/groupbuys/:id", controllers.GetGroupBuy)
		auth.POST("/groupbuys/:id/join", controllers.JoinGroupBuy)

		auth.POST("/orders", controllers.CreateOrder)
		auth.GET("/orders", controllers.GetOrders)
		auth.GET("/orders/:id", controllers.GetOrder)
		auth.DELETE("/orders/:id", controllers.DeleteOrder)
	}

	log.Println("Server http://localhost:8080")
	r.Run()
}
