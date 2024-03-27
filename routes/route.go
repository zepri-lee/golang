package routes

import (
	"gin-gonic-gorm/config"
	"gin-gonic-gorm/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	app.Use(corsConfig())
	app.Static(config.STATIC_ROUTE, config.STATIC_DIR)

	api := app.Group("/api")
	{
		api.GET("/stock", controllers.GetStock)
		api.GET("/stockPaging", controllers.GetStockPaging)
		api.POST("/stock", controllers.AddStock)
		api.POST("/sale", controllers.AddSale)
		api.DELETE("/stockAll", controllers.DeleteAllStock)
		api.DELETE("/stockById", controllers.DeleteStockById)
		api.DELETE("/stockById2/:productId", controllers.DeleteStockById2)

		api.POST("/file", controllers.HandleUploadFile)
		api.DELETE("/file/:fileName", controllers.HadndleRemoveFile)
	}
}

func corsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Accept-Language", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	})
}
