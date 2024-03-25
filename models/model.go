package models

import (
	"time"
)

type Stock struct {
	ProductID     uint   `json:"productId" gorm:"column:PRODUCT_ID"`     // 상품의 고유 식별자
	ProductName   string `json:"productName" gorm:"column:PRODUCT_NAME"` // 상품명
	StockQuantity int    `json:"quantity" gorm:"column:STOCK_QUANTITY"`  // 해당 상품의 수량
	StockLocation string `json:"location" gorm:"column:STOCK_LOCATION"`  // 상품이 위치한 장소
}

type Sale struct {
	ProductID uint      `json:"productId" gorm:"column:PRODUCT_ID"` // 상품의 고유 식별자
	SaleCount int       `json:"count" gorm:"column:SALE_COUNT"`     // 판매 수량
	SaleDate  time.Time `json:"date" gorm:"column:SALE_DATE"`       // 판매 시점
}
