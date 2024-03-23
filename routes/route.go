package routes

import (
	controller "gin-gonic-gorm/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app

	route.GET("/user", controller.GetAllUser)

	route.GET("/book", controller.GetAllBook)
}
