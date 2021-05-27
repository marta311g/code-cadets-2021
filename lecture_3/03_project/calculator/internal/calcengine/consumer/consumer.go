package consumer

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Consumer struct {
	betConsumer         BetConsumer
	eventUpdateConsumer EventUpdateConsumer
}

func New(eventUpdateConsumer EventUpdateConsumer, betConsumer BetConsumer) *Consumer {
	return &Consumer{
		betConsumer: betConsumer,
		eventUpdateConsumer: eventUpdateConsumer,
	}
}

func (c *Consumer) ConsumeEventUpdates(ctx context.Context) (<-chan rabbitmqmodels.EventUpdate, error) {
	return c.eventUpdateConsumer.Consume(ctx)
}

func (c *Consumer) ConsumeBets(ctx context.Context) (<-chan rabbitmqmodels.Bet, error) {
	return c.betConsumer.Consume(ctx)
}