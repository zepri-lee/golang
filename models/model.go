package models

import "time"

type Stock struct {
	PRODUCTID      uint   `json:"productID"`   // 상품의 고유 식별자
	PRODUCTNAME    string `json:"productName"` // 상품명
	STOCK_QUANTITY int    `json:"quantity"`    // 해당 상품의 수량
	STOCK_LOCATION string `json:"location"`    // 상품이 위치한 장소
}

type Sale struct {
	PRODUCTID  uint      `json:"productID"` // 상품의 고유 식별자
	SALE_COUNT int       `json:"count"`     // 판매 수량
	SALE_DATE  time.Time `json:"date"`      // 판매 시점
}
