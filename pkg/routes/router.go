package routes

import (
	"assignment2/pkg/controllers"

	"github.com/gin-gonic/gin"
)

type routers struct {
	router       *gin.Engine
	orderController *controllers.OrderController
}

func NewRouter(router *gin.Engine, orderController *controllers.OrderController) *routers {
	return &routers{router, orderController}
}

func (r *routers) Start(port string) {
	r.router.POST("/order", r.orderController.CreateNewOrder)

	// r.router.GET("/orders", r.orderController.GetOrders)
	r.router.Run(port)
}