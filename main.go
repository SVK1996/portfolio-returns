package main

import (
	"fmt"
	"time"

	"github.com/SVK1996/portfolio-returns/models"
	"github.com/SVK1996/portfolio-returns/services"
)

func main() {
	// Sample data
	portfolios := []models.Portfolio{
		{ID: "trader1", InitialCap: 1000000, Orders: []models.Order{
			{PortfolioID: "trader1", StockID: "stock1", TransactionType: "BUY", Quantity: 10, Price: 100.0, Timestamp: time.Now()},
			{PortfolioID: "trader1", StockID: "stock1", TransactionType: "SELL", Quantity: 10, Price: 110.0, Timestamp: time.Now().Add(24 * time.Hour)},
		}},
		{ID: "trader2", InitialCap: 1000000, Orders: []models.Order{
			{PortfolioID: "trader2", StockID: "stock2", TransactionType: "BUY", Quantity: 20, Price: 50.0, Timestamp: time.Now()},
			{PortfolioID: "trader2", StockID: "stock2", TransactionType: "SELL", Quantity: 20, Price: 55.0, Timestamp: time.Now().Add(24 * time.Hour)},
		}},
	}

	returns, err := services.CalculateReturns(portfolios)
	if err != nil {
		fmt.Printf("Error calculating returns: %v\n", err)
		return
	}
	fmt.Printf("Returns: %+v\n", returns)
}
