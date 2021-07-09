package calcengine

import (
	"context"
	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Handler interface {
	HandleEventUpdates(ctx context.Context, eventUpdates <-chan rabbitmqmodels.EventUpdate) <-chan rabbitmqmodels.BetCalculated
	HandleBets(ctx context.Context, bets <-chan rabbitmqmodels.Bet) error
}
