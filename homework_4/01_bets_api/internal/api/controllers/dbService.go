package controllers

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/api/controllers/models"
)

// DbService implements database related functions.
type DbService interface {
	GetBetByID(id string) (models.BetResponseDto, error)
	GetBetsByUser(customerId string) ([]models.BetResponseDto, error)
	GetBetsByStatus(status string) ([]models.BetResponseDto, error)
}