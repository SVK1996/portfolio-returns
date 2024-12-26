package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/SVK1996/portfolio-returns/models"
)

func CalculatePortfolioValue(portfolio models.Portfolio, timestamp time.Time) (float64, error) {
	if len(portfolio.Orders) == 0 {
		return 0.0, errors.New("no orders in portfolio")
	}

	// Map to keep track of the quantity of each stock
	stockQuantities := make(map[string]int)

	// Calculate the quantity of each stock held in the portfolio
	for _, order := range portfolio.Orders {
		if order.Timestamp.After(timestamp) {
			continue
		}
		switch order.TransactionType {
		case "BUY":
			stockQuantities[order.StockID] += order.Quantity
		case "SELL":
			stockQuantities[order.StockID] -= order.Quantity
		case "SHORT":
			stockQuantities[order.StockID] -= order.Quantity
		case "COVER":
			stockQuantities[order.StockID] += order.Quantity
		default:
			return 0.0, errors.New("invalid transaction type")
		}
	}

	// Calculate the total value of the portfolio
	totalValue := portfolio.InitialCap
	for stockID, quantity := range stockQuantities {
		currentPrice, err := getCurrentStockPrice(stockID, timestamp)
		if err != nil {
			return 0.0, err
		}
		totalValue += currentPrice * float64(quantity)
	}

	return totalValue, nil
}

// Function to get the current price of a stock
func getCurrentStockPrice(stockID string, timestamp time.Time) (float64, error) {
	// Mock API endpoint
	apiURL := fmt.Sprintf("https://mock-api.com/stock-price?stockID=%s&timestamp=%s", stockID, timestamp.Format(time.RFC3339))

	resp, err := http.Get(apiURL)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0.0, errors.New("failed to fetch stock price")
	}

	// Assume the API returns a JSON response with the stock price
	type StockPriceResponse struct {
		Price float64 `json:"price"`
	}

	var priceResponse StockPriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&priceResponse); err != nil {
		return 0.0, err
	}

	return priceResponse.Price, nil
}
