package services

// BetReceivedPublisher handles bet received queue publishing.
type BetReceivedPublisher interface {
	Publish(customerId, selectionId string, selectionCoefficient, payment float64) error
}
