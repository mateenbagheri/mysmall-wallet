package models

type Transaction struct {
	TransactionID   int64   `json:"reference_id"`
	UserID          int     `json:"user_id" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
	TransactionDate string  `json:"transaction_date"`
}

type Reference struct {
	TransactionID int64 `json:"reference_id"`
}
