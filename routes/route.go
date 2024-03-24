package routes

import (
	controller "gin-gonic-gorm/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	api := app.Group("/api")
	{
		api.GET("/stock", controller.GetStock)
		api.POST("/stock", controller.AddStock)
		api.DELETE("/stock", controller.DeleteAllStock)
		api.POST("/sale", controller.AddSale)
	}
}
