package controllers

import (
	"errors"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/******************************************************************************************************
* 재고정보 조회
* - location은 기본조건이고 나머지는 filter조건
* - Quantity FROM, TO 조회를 할 경우 FROM, TO 조건 모두 존재해야 함
*******************************************************************************************************/
func GetStock(ctx *gin.Context) {
	//	stocks := new([]models.Stock)
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
		futureQuery = futureQuery.Where("PRODUCT_NAME  = ?", productName)
	}
	if frQty != "" && toQty != "" {
		futureQuery = futureQuery.Where("STOCK_QUANTITY BETWEEN ? AND ?", frQty, toQty)
	}
	// 재고정보 체크
	//	if err := futureQuery.First(&stocks).Error; err != nil {
	//		// 데이터 미존재
	//		if errors.Is(err, gorm.ErrRecordNotFound) {
	//			ctx.JSON(http.StatusOK, gin.H{
	//				"message": "재고정보 미존재1",
	//			})
	//		}
	//
	//		return
	//	}

	// 데이터 조회
	if err := futureQuery.Scan(&stocks).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	// 데이터 존재여부 체크
	if len(stocks) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "재고정보 미존재2",
		})

		return
	}

	// 정상조회
	ctx.JSON(http.StatusOK, gin.H{
		"data": stocks,
	})
}

/******************************************************************************************************
* 재고정보 등록
* 재고가 존재하면 업데이트(QUANTITY 증가), 존재하지 않으면 인서트
*******************************************************************************************************/
func AddStock(ctx *gin.Context) {

	var stocks []models.Stock
	var stock_count int = 0

	// 파라미터 바인딩
	if err := ctx.ShouldBindJSON(&stocks); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	// 존재하면 업데이트(QUANTITY 증가), 존재하지 않으면 인서트
	for _, stock := range stocks {
		if err := database.Instance.Select("COUNT(*) AS STOCK_COUNT").Table("STOCK").
			Where("PRODUCT_ID = ?", stock.PRODUCT_ID).Find(&stock_count).Error; err != nil {

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})

			return
		}

		if stock_count == 0 {
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

			err := database.Instance.Exec(query, stock.STOCK_QUANTITY, stock.PRODUCT_ID).Error
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

	ctx.JSON(http.StatusOK, gin.H{
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

/******************************************************************************************************
* 판매정보 등록
* 1. 재고체크 => 미존재시 오류, 재고부족
* 2. 판매등록
* 3. 재고차감
*******************************************************************************************************/
func AddSale(ctx *gin.Context) {

	var sale models.Sale
	var stock_qunatity int = 0
	var stock_count int = 0

	tx := database.Instance.Begin()
	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "transaction start failed",
		})
		return
	}

	// 파라미터 바인딩
	if err := ctx.ShouldBindJSON(&sale); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	// 판매일자
	sale.SALE_DATE = time.Now()

	// 재고체크 => 미존재
	if err := tx.Select("COUNT(*)").Table("STOCK").
		Where("PRODUCT_ID = ?", sale.PRODUCT_ID).Find(&stock_count).Error; err != nil {

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	if stock_count == 0 {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "상품이 재고정보에 등록되어 있지 않습니다.",
		})

		return
	}

	// 재고체크 => 재고부족
	if err := tx.Select("ISNULL(STOCK_QUANTITY, 0)").Table("STOCK").
		Where("PRODUCT_ID = ?", sale.PRODUCT_ID).Find(&stock_qunatity).Error; err != nil {

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	if stock_qunatity <= sale.SALE_COUNT {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "상품의 재고정보가 부족합니다.",
		})

		return
	}

	// 재고 차감
	query := `UPDATE STOCK
			 SET STOCK_QUANTITY = STOCK_QUANTITY - ?
			 WHERE PRODUCT_ID = ?`

	if err := tx.Exec(query, sale.SALE_COUNT, sale.PRODUCT_ID).Error; err != nil {
		tx.Rollback()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 판매등록
	if err := tx.Table("SALE").Create(&sale).Error; err != nil {
		tx.Rollback()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// ID로 재고 삭제 (Get request 방식)
func DeleteStockById(ctx *gin.Context) {
	//	productId := ctx.Param("productId")
	productId := ctx.Query("productId")
	var stocks models.Stock

	// 재고 존재유무 체크
	errFirst := database.Instance.Table("STOCK").Where("PRODUCT_ID = ?", productId).First(&stocks).Error

	if errors.Is(errFirst, gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "재고 미존재",
		})
		return
	}
	if errFirst != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errFirst.Error(),
		})
		return
	}

	// 삭제
	if err := database.Instance.Table("STOCK").Unscoped().Where("PRODUCT_ID = ?", productId).Delete(&models.Stock{}).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// ID로 재고 삭제 (uri 셋팅방식)
func DeleteStockById2(ctx *gin.Context) {
	productId := ctx.Param("productId")
	//	productId := ctx.Query("productId")
	var stocks models.Stock

	// 재고 존재유무 체크
	errFirst := database.Instance.Table("STOCK").Where("PRODUCT_ID = ?", productId).First(&stocks).Error

	if errors.Is(errFirst, gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "재고 미존재",
		})
		return
	}
	if errFirst != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": errFirst.Error(),
		})
		return
	}

	// 삭제
	if err := database.Instance.Table("STOCK").Unscoped().Where("PRODUCT_ID = ?", productId).Delete(&models.Stock{}).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// 테스트용 전제 삭제
func DeleteAllStock(ctx *gin.Context) {

	if err := database.Instance.Exec("DELETE FROM STOCK").Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}
}
