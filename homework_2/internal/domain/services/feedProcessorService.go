package services

import (
	"context"
	"log"

	"github.com/pkg/errors"

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
			case msg, ok := <- updates[counter % len(updates)]:
				if !ok {
					return nil
				}
				msg.Coefficient *= 2
				source <- msg
		}
		counter += 1
	}

	return errors.New("feed processor service")
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
