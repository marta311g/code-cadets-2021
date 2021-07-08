package calcengine

import (
	"context"
	"log"
)


type Engine struct {
	consumer  Consumer
	handler   Handler
	publisher Publisher
}

func New(consumer Consumer, handler Handler, publisher Publisher) *Engine {
	return &Engine{
		consumer:  consumer,
		handler:   handler,
		publisher: publisher,
	}
}

func (e *Engine) Start(ctx context.Context) {
	err := e.processEventUpdates(ctx)
	if err != nil {
		log.Println("Engine failed to process event updates:", err)
		return
	}

	err = e.processBets(ctx)
	if err != nil {
		log.Println("Engine failed to process bets:", err)
		return
	}

	<-ctx.Done()
}

func (e *Engine) processEventUpdates(ctx context.Context) error {
	consumedEventUpdates, err := e.consumer.ConsumeEventUpdates(ctx)
	if err != nil {
		return err
	}

	newBets := e.handler.HandleEventUpdates(ctx, consumedEventUpdates)
	e.publisher.PublishBetsCalculated(ctx, newBets)

	return nil
}

func (e *Engine) processBets(ctx context.Context) error {
	consumedBets, err := e.consumer.ConsumeBets(ctx)
	if err != nil {
		return err
	}

	err = e.handler.HandleBets(ctx, consumedBets)

	return err
}