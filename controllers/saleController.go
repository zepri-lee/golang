package controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetStock(ctx *gin.Context) {
	//stocks := new([]models.Stock)
	var stocks []models.Stock
	//	location := ctx.Param("location") // http://~~~~/:location
	location := ctx.Query("location")
	productName := ctx.Query("productName")
	frQty := ctx.Query("frQty")
	toQty := ctx.Query("toQty")

	// location 제약 조건 체크
	if location == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "STOCK Location을 입력바랍니다.",
		})

		return
	}

	// frQty, toQty 제약 조건 체크
	if (frQty == "" && toQty != "") ||
		(frQty != "" && toQty == "") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "STOCK QUANTITY의 FROM, TO 모두 입력바랍니다.",
		})

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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	// 데이터 존재여부 체크
	if len(stocks) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "데이터 미존재",
		})

		return
	}

	// 정상조회
	ctx.JSON(http.StatusOK, gin.H{
		"data": stocks,
	})
}

func AddStock(ctx *gin.Context) {
	var stocks []models.Stock

	type STOCK_COUNT struct {
		STOCK_COUNT uint
	}

	var stock_count STOCK_COUNT

	if err := ctx.ShouldBindJSON(&stocks); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	// 존재하면 업데이트(QUANTITY 증가), 존재하지 않으면 인서트
	for _, stock := range stocks {
		if err := database.Instance.Select("COUNT(*) AS STOCK_COUNT").Table("STOCK").
			Where("PRODUCT_ID = ?", stock.Product_ID).Find(&stock_count).Error; err != nil {

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})

			return
		}

		if stock_count.STOCK_COUNT == 0 {
			if err := database.Instance.Table("STOCK").Create(&stock).Error; err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})

				database.Instance.Rollback()
				return
			}
		} else {
			query := `UPDATE STOCK
					 SET STOCK_QUANTITY = STOCK_QUANTITY + ?
					 WHERE PRODUCT_ID = ?`

			err := database.Instance.Exec(query, stock.Stock_Quantity, stock.Product_ID).Error
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})

				database.Instance.Rollback()
				return
			}
			/* 			if err := database.Instance.Table("STOCK").Update("STOCK_QUANTITY", gorm.Expr("STOCK_QUANTITY + ?", stock.Stock_Quantity)).
				Where("PRODUCT_ID = ?", stock.Product_ID).Error; err != nil {

				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})

				database.Instance.Rollback()
				return
			} */
		}
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "OK",
	})
	/*
		if err := database.Instance.Table("STOCK").CreateInBatches(&stocks, 10).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "OK",
			})
		}
	*/
}

// 테스트용 전제 삭제
func DeleteAllStock(ctx *gin.Context) {

	if err := database.Instance.Exec("DELETE FROM STOCK").Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "OK",
		})
	}
}

func AddSale(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "AddSale",
	})
}
