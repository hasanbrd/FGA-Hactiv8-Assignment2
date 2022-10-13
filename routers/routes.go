package routers

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreteOrder)
	router.GET("/orders", controllers.GetAllOrders)
	router.GET("/orders/:orderID", controllers.GetOrderByID)
	router.PUT("/orders/:orderID", controllers.UpdateOrder)
	router.DELETE("orders/:orderID", controllers.DeleteOrder)

	return router
}
