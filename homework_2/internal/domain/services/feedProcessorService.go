package services

import (
	"context"
	"log"

	"code-cadets-2021/homework_2/internal/domain/models"
)

type FeedProcessorService struct {
	feeds  []Feed
	queue Queue
}

func NewFeedProcessorService(
	feeds []Feed,
	queue Queue,
) *FeedProcessorService {
	return &FeedProcessorService{
		feeds:  feeds,
		queue: queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	source := f.queue.GetSource()
	var updates []chan models.Odd

	for _, feed := range f.feeds {
		updates = append(updates, feed.GetUpdates())
	}

	defer close(source)
	defer log.Printf("shutting down %s", f)

	counter := 0
	for {
		select {
			case msg1, ok := <- updates[0]:
				if !ok {
					return nil
				}
				msg1.Coefficient *= 2
				source <- msg1
			case msg2, ok := <- updates[1]:
				if !ok {
					return nil
				}
				msg2.Coefficient *= 2
				source <- msg2
			case <- ctx.Done():
				log.Println("processor service: context canceled")
				return nil
		}
		counter += 1
	}
}

func (f *FeedProcessorService) String() string {
	return "feed processor service"
}

type Feed interface {
	GetUpdates() chan models.Odd
}

type Queue interface {
	GetSource() chan models.Odd
}
