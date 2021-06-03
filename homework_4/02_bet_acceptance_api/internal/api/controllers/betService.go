package controllers

// BetService implements bet related functions.
type BetService interface {
	InsertBet(customerId string, selectionId string, selectionCoefficient float64, payment float64) error
}
