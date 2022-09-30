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
	r.router.GET("/orders", r.orderController.GetAllOrdersWithItems)
	r.router.DELETE("/orders/:orderId", r.orderController.DeleteOrder)
	r.router.PUT("/orders/:orderId", r.orderController.UpdateOrder)
	r.router.Run(port)
}