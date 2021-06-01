package controllers

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/api/controllers/models"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/infrastructure/sqlite"
)

// EventService implements event related functions.
type DbService interface {
	GetBetByID(id string, dbExecutor sqlite.DatabaseExecutor) (models.BetResponseDto, error)
	GetBetsByUser(customerId string) ([]models.BetResponseDto, error)
	GetBetsByStatus(status string) ([]models.BetResponseDto, error)
}
