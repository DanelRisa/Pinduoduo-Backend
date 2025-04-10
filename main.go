package main

import (
	"log"
	"pinduoduo-back/controllers"
	"pinduoduo-back/database"


	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.Migrate()

	r := gin.Default()

	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProduct)
	r.POST("/products", controllers.CreateProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	r.POST("/groupbuys", controllers.CreateGroupBuy)
	r.GET("/groupbuys", controllers.GetGroupBuys)
	r.GET("/groupbuys/:id", controllers.GetGroupBuy)
	r.POST("/groupbuys/:id/join", controllers.JoinGroupBuy)

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrders)
	r.GET("/orders/:id", controllers.GetOrder)
	r.DELETE("/orders/:id", controllers.DeleteOrder)

	log.Println("Server http://localhost:8080")
	r.Run()
}
