package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/pkg/errors"

	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)


type EventUpdateConsumer struct {
	channel Channel
	config  ConsumerConfig
}


func NewEventUpdateConsumer(channel Channel, config ConsumerConfig) (*EventUpdateConsumer, error) {
	_, err := channel.QueueDeclare(
		config.Queue,
		config.DeclareDurable,
		config.DeclareAutoDelete,
		config.DeclareExclusive,
		config.DeclareNoWait,
		config.DeclareArgs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "bet received consumer initialization failed")
	}

	return &EventUpdateConsumer{
		channel: channel,
		config:  config,
	}, nil
}


func (c *EventUpdateConsumer) Consume(ctx context.Context) (<-chan models.EventUpdate, error) {
	msgs, err := c.channel.Consume(
		c.config.Queue,
		c.config.ConsumerName,
		c.config.AutoAck,
		c.config.Exclusive,
		c.config.NoLocal,
		c.config.NoWait,
		c.config.Args,
	)
	if err != nil {
		return nil, errors.Wrap(err, "bet received consumer failed to consume messages")
	}

	eventUpdates := make(chan models.EventUpdate)
	go func() {
		defer close(eventUpdates)
		for msg := range msgs {
			var eventUpdate models.EventUpdate
			err := json.Unmarshal(msg.Body, &eventUpdate)
			if err != nil {
				log.Println("Failed to unmarshal event update message", msg.Body)
			}

			// Once context is cancelled, stop consuming messages.
			select {
			case eventUpdates <- eventUpdate:
			case <-ctx.Done():
				return
			}
		}
	}()

	return eventUpdates, nil
}
