package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"pinduoduo-back/database"
	"pinduoduo-back/order-service/controllers"
	"pinduoduo-back/order-service/middleware"

	// "github.com/go-resty/resty/v2"
    // "net/http"
)

func main() {
	database.Connect()

	r := gin.Default()

	auth := r.Group("/")
	auth.Use(middleware.LoggingMiddleware())

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

	log.Println("Order service running on http://localhost:8082")
	r.Run(":8082")
}
