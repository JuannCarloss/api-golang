package routers

import (
	"example/payments/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	r.GET("/payment", controllers.FindAll)
	r.POST("/payment", controllers.PostPayment)

	return r
}
