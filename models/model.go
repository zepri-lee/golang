package models

import (
	"time"
)

type Stock struct {
	Product_ID     uint   `json:"productId"`   // 상품의 고유 식별자
	Product_Name   string `json:"productName"` // 상품명
	Stock_Quantity int    `json:"quantity"`    // 해당 상품의 수량
	Stock_Location string `json:"location"`    // 상품이 위치한 장소
}

type Sale struct {
	Product_ID uint      `json:"productId"` // 상품의 고유 식별자
	Sale_Count int       `json:"count"`     // 판매 수량
	Sale_Date  time.Time `json:"date"`      // 판매 시점
}
