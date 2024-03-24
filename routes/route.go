package routes

import (
	"gin-gonic-gorm/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	api := app.Group("/api")
	{
		api.GET("/stock", controllers.GetStock)
		api.POST("/stock", controllers.AddStock)
		api.DELETE("/stock", controllers.DeleteAllStock)
		api.POST("/sale", controllers.AddSale)
	}
}
