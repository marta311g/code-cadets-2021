package services

import (
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
	"context"
)

type FeedProcessorService struct {
	feed FeedInterface
	queue QueueInterface
}

func NewFeedProcessorService(feed FeedInterface, queue QueueInterface) *FeedProcessorService {
	// it should receive "Feed" & "Queue" interfaces through constructor
	return &FeedProcessorService{
		feed: feed,
		queue: queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	// initially:
	// - get updates channel from feed interface
	updates := f.feed.GetUpdates()
	// - get source channel from queue interface
	source := f.queue.GetSource()
	//
	// repeatedly:
	// - range over updates channel
	for update := range updates {
		// - multiply each odd with 2
		update.Coefficient = update.Coefficient*2
		// - send it to source channel
		source <- update
	}
	//
	// finally:
	// - when updates channel is closed, exit
	// - when exiting, close source channel
	close(source)
	return nil
}

// define feed interface here
type FeedInterface interface {
	GetUpdates() chan models.Odd
}

// define queue interface here
type QueueInterface interface {
	GetSource() chan models.Odd
}
