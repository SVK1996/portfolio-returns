package utils

import "time"

func IsMarketDay(timestamp time.Time) bool {
	// Check if the timestamp falls on a weekday (Monday to Friday)
	weekday := timestamp.Weekday()
	return weekday >= time.Monday && weekday <= time.Friday
}
