package models

type Portfolio struct {
	ID         string
	InitialCap float64
	Orders     []Order
}
