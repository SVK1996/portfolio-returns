package models

import "time"

type Order struct {
	PortfolioID     string
	StockID         string
	TransactionType string
	Quantity        int
	Price           float64
	Timestamp       time.Time
}
