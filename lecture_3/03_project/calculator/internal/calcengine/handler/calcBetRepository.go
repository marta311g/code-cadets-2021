package handler

import (
	"context"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
)

type CalcBetRepository interface {
	InsertCalcBet(ctx context.Context, bet domainmodels.Bet) error
	UpdateCalcBet(ctx context.Context, bet domainmodels.Bet) error
	CalcBetWithIDExists(ctx context.Context, id string) (bool, error)
	GetBetsByEventID(ctx context.Context, eventId string) ([]domainmodels.Bet, bool, error)
}