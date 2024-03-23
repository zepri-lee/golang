package controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetStock(ctx *gin.Context) {
	stocks := new([]models.Stock)
	//	location := ctx.Param("location")
	location := ctx.Query("location")
	productName := ctx.Query("productName")
	frQty := ctx.Query("frQty")
	toQty := ctx.Query("toQty")

	// frQty, toQty 두개 모두 존재해야 함
	if location == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "STOCK Location을 입력바랍니다.",
		})
		ctx.Abort()
		return
	}

	// frQty, toQty 두개 모두 존재해야 함
	if (frQty == "" && toQty != "") ||
		(frQty != "" && toQty == "") {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "STOCK QUANTITY의 FROM, TO 모두 입력바랍니다.",
		})
		ctx.Abort()
		return
	}

	// Define the common part of the query
	commonQuery := func() *gorm.DB {
		return database.Instance.
			Table("STOCK").
			Where("STOCK_LOCATION = ?", location)
	}

	// 조건존재여부에 따라 쿼리 조합
	// location은 기본조건이고 나머지는 filter조건
	futureQuery := commonQuery()
	if productName != "" {
		futureQuery = futureQuery.Where("PRODUCTNAME  = ?", productName)
	}
	if frQty != "" && toQty != "" {
		futureQuery = futureQuery.Where("STOCK_QUANTITY BETWEEN ? AND ?", frQty, toQty)
	}
	if err := futureQuery.Scan(&stocks).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		ctx.Abort()
		return
	}

	// 정상조회
	ctx.JSON(200, gin.H{
		"data": stocks,
	})
}

func AddStock(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "AddStock",
	})
}

func AddSale(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "AddSale",
	})
}
