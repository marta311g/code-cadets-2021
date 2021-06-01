package services

import (
	"context"
	"github.com/pkg/errors"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/api/controllers/models"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/domain/mappers"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/infrastructure/sqlite"
)

// EventService implements event related functions.
type BetService struct {
	dbExecutor sqlite.DatabaseExecutor
}

// NewEventService creates a new instance of EventService.
func NewBetService(dbExecutor sqlite.DatabaseExecutor) *BetService {
	return &BetService{
		dbExecutor: dbExecutor,
	}
}

func newBetMapper() *mappers.BetMapper {
	return mappers.NewBetMapper()
}

// UpdateEvent sends event update message to the queues.
func (e BetService) GetBetByID(id string, dbExecutor sqlite.DatabaseExecutor) (models.BetResponseDto, error) {
	betMapper := newBetMapper()
	betRepository := sqlite.NewBetRepository(dbExecutor, betMapper)
	bet, exists, err := betRepository.GetBetByID(context.Background(), id)

	if err != nil {
		return models.BetResponseDto{}, errors.WithMessage(err, "error")
	}
	if !exists {
		return models.BetResponseDto{}, errors.WithMessage(err, "doesn't exist")
	}

	return bet, nil
}

func (e BetService) GetBetsByUser(userId string) ([]models.BetResponseDto, error) {
	betMapper := newBetMapper()
	betRepository := sqlite.NewBetRepository(e.dbExecutor, betMapper)
	bets, exists, err := betRepository.GetBetsByUserId(context.Background(), userId)

	if err != nil {
		return []models.BetResponseDto{}, errors.WithMessage(err, "error")
	}
	if !exists {
		return []models.BetResponseDto{}, errors.WithMessage(err, "doesn't exist")
	}

	return bets, nil
}

func (e BetService) GetBetsByStatus(status string) ([]models.BetResponseDto, error) {
	return []models.BetResponseDto{}, nil
}
