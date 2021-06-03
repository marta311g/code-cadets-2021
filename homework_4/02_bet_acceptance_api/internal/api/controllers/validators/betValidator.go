package validators

import "github.com/superbet-group/code-cadets-2021/homework_4/02_bet_acceptance_api/internal/api/controllers/models"

// BetValidator validates bet requests.
type BetValidator struct{}

// NewBetValidator creates a new instance of BetValidator.
func NewBetValidator() *BetValidator {
	return &BetValidator{}
}

// BetIsValid checks if bet is valid.
func (e *BetValidator) BetIsValid(betInsertRequestDto models.BetInsertRequestDto) bool {
	if betInsertRequestDto.CustomerId != "" && betInsertRequestDto.SelectionId != "" && (betInsertRequestDto.SelectionCoefficient != 0.0 && betInsertRequestDto.SelectionCoefficient <= 10.0) && (betInsertRequestDto.Payment >= 2.0 && betInsertRequestDto.Payment <= 100.0) {
		return true
	}
	return false
}
