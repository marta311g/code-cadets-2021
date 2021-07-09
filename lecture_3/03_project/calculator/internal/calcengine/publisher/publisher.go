package publisher

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Publisher struct {
	betCalculatedPublisher BetCalculatedPublisher
}

func New(betCalculatedPublisher BetCalculatedPublisher) *Publisher {
	return &Publisher{
		betCalculatedPublisher: betCalculatedPublisher,
	}
}

func(p *Publisher) PublishBetsCalculated(ctx context.Context, betsCalculated <- chan rabbitmqmodels.BetCalculated) {
	p.betCalculatedPublisher.Publish(ctx, betsCalculated)
}