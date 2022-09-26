package main

import (
	"assignment2/pkg/config"
	"assignment2/pkg/controllers"
	"assignment2/pkg/repositories"
	"assignment2/pkg/routes"
	"assignment2/pkg/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.StartDB()
	if err != nil {
		panic(err)
	}
	orderRepo := repositories.NewOrderRepo(db)
	orderService := services.NewOrderService(orderRepo)

	itemRepo := repositories.NewItemRepo(db)
	itemService := services.NewItemService(itemRepo)
	orderController := controllers.NewOrderController(orderService, itemService)


	router := gin.Default()

	app := routes.NewRouter(router, orderController)
	app.Start(":4000")
	
}