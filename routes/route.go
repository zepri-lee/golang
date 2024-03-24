package routes

import (
	"gin-gonic-gorm/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	app.Use(corsConfig())

	api := app.Group("/api")
	{
		api.GET("/stock", controllers.GetStock)
		api.POST("/stock", controllers.AddStock)
		api.DELETE("/stock", controllers.DeleteAllStock)
		api.POST("/sale", controllers.AddSale)
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
