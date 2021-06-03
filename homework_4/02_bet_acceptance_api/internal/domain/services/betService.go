package services

// BetService implements bet related functions.
type BetService struct {
	betReceivedPublisher BetReceivedPublisher
}

// NewBetService creates a new instance of BetService.
func NewBetService(betReceivedPublisher BetReceivedPublisher) *BetService {
	return &BetService{
		betReceivedPublisher: betReceivedPublisher,
	}
}

// InsertBet sends new bet to the queues.
func (e BetService) InsertBet(customerId string, selectionId string, selectionCoefficient float64, payment float64) error {
	return e.betReceivedPublisher.Publish(customerId, selectionId, selectionCoefficient, payment)
}
