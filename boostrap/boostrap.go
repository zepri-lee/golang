package boostrap

import (
	"gin-gonic-gorm/config"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"gin-gonic-gorm/routes"

	"github.com/gin-gonic/gin"
)

func BoostrapApp() {
	database.Connect()
	app := gin.Default()
	// 테이블 생성
	database.Instance.AutoMigrate(&models.Stock{})
	database.Instance.AutoMigrate(&models.Sale{})

	routes.InitRoute(app)

	app.Run(config.PORT)
}
