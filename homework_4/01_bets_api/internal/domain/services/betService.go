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
	repository sqlite.BetRepository
}

// NewEventService creates a new instance of EventService.
func NewBetService(repository sqlite.BetRepository) *BetService {
	return &BetService{
		repository: repository,
	}
}

func newBetMapper() *mappers.BetMapper {
	return mappers.NewBetMapper()
}

// UpdateEvent sends event update message to the queues.
func (e BetService) GetBetByID(id string) (models.BetResponseDto, error) {
	bet, exists, err := e.repository.GetBetByID(context.Background(), id)

	if err != nil {
		return models.BetResponseDto{}, err
	}
	if !exists {
		return models.BetResponseDto{}, errors.WithMessage(err, "bet "+id+" doesn't exist")
	}

	return bet, nil
}

func (e BetService) GetBetsByUser(userId string) ([]models.BetResponseDto, error) {
	bets, exists, err := e.repository.GetBetsByUserId(context.Background(), userId)

	if err != nil {
		return []models.BetResponseDto{}, err
	}
	if !exists {
		return []models.BetResponseDto{}, errors.WithMessage(err, "no bets for this user "+userId)
	}

	return bets, nil
}

func (e BetService) GetBetsByStatus(status string) ([]models.BetResponseDto, error) {
	bets, exists, err := e.repository.GetBetsByStatus(context.Background(), status)

	if err != nil {
		return []models.BetResponseDto{}, err
	}
	if !exists {
		return []models.BetResponseDto{}, errors.WithMessage(err, "no bets with status "+status)
	}

	return bets, nil
}
