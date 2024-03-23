package boostrap

import (
	"gin-gonic-gorm/config"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/routes"

	"github.com/gin-gonic/gin"
)

func BoostrapApp() {
	database.Connect()
	app := gin.Default()

	routes.InitRoute(app)

	app.Run(config.PORT)
}
