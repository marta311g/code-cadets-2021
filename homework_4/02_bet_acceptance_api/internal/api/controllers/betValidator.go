package controllers

import "github.com/superbet-group/code-cadets-2021/homework_4/02_bet_acceptance_api/internal/api/controllers/models"

// BetValidator validates bet requests.
type BetValidator interface {
	BetIsValid(betInsertRequestDto models.BetInsertRequestDto) bool
}
