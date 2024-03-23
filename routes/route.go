package routes

import (
	controller "gin-gonic-gorm/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	strApi := "/api"
	route := app

	route.GET(strApi+"/stock", controller.GetStock)
	route.POST(strApi+"/stock", controller.AddStock)
	route.DELETE(strApi+"/stock", controller.DeleteAllStock)
	route.POST(strApi+"/sale", controller.AddSale)
}
