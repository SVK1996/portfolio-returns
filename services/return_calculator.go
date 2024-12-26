package services

import (
	"errors"
	"sort"

	"github.com/SVK1996/portfolio-returns/models"
)

func CalculateReturns(portfolios []models.Portfolio) (map[string][]models.Return, error) {
	if len(portfolios) == 0 {
		return nil, errors.New("no portfolios provided")
	}

	returns := make(map[string][]models.Return)
	var err error
	returns["daily"], err = calculateDailyReturns(portfolios)
	if err != nil {
		return nil, err
	}
	returns["monthly"], err = calculateMonthlyReturns(portfolios)
	if err != nil {
		return nil, err
	}
	returns["yearly"], err = calculateYearlyReturns(portfolios)
	if err != nil {
		return nil, err
	}
	returns["lifetime"], err = calculateLifetimeReturns(portfolios)
	if err != nil {
		return nil, err
	}
	return returns, nil
}

func calculateDailyReturns(portfolios []models.Portfolio) ([]models.Return, error) {
	return calculateReturnsForPeriod(portfolios, 1)
}

func calculateMonthlyReturns(portfolios []models.Portfolio) ([]models.Return, error) {
	return calculateReturnsForPeriod(portfolios, 22)
}

func calculateYearlyReturns(portfolios []models.Portfolio) ([]models.Return, error) {
	return calculateReturnsForPeriod(portfolios, 264)
}

func calculateLifetimeReturns(portfolios []models.Portfolio) ([]models.Return, error) {
	return calculateReturnsForPeriod(portfolios, -1)
}

func calculateReturnsForPeriod(portfolios []models.Portfolio, period int) ([]models.Return, error) {
	var returns []models.Return
	for _, portfolio := range portfolios {
		returnValue, err := calculatePortfolioReturn(portfolio, period)
		if err != nil {
			return nil, err
		}
		returns = append(returns, models.Return{TraderID: portfolio.ID, Return: returnValue})
	}
	sort.Slice(returns, func(i, j int) bool {
		return returns[i].Return > returns[j].Return
	})
	if len(returns) > 50 {
		return returns[:50], nil
	}
	return returns, nil
}

func calculatePortfolioReturn(portfolio models.Portfolio, period int) (float64, error) {
	if len(portfolio.Orders) == 0 {
		return 0.0, errors.New("no orders in portfolio")
	}

	// Implement the logic to calculate the return based on the orders and period
	initialValue := portfolio.InitialCap
	finalValue := initialValue

	for _, order := range portfolio.Orders {
		switch order.TransactionType {
		case "BUY":
			finalValue -= order.Price * float64(order.Quantity)
		case "SELL":
			finalValue += order.Price * float64(order.Quantity)
		case "SHORT":
			finalValue += order.Price * float64(order.Quantity)
		case "COVER":
			finalValue -= order.Price * float64(order.Quantity)
		default:
			return 0.0, errors.New("invalid transaction type")
		}
	}

	return ((finalValue - initialValue) / initialValue) * 100, nil
}
