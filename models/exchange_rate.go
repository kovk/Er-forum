package models

import "time"

type ExchangeRate struct {
	ID           uint      `gorm:"primary_key" json:"_id"`
	FromCurrency string    `json:"fromCurrency" binding:"required"`
	Tocurrency   string    `json:"toCurrency" binding:"required"`
	Rate         float64   `json:"rate" binding:"required"`
	Date         time.Time `json:"date"`
}
