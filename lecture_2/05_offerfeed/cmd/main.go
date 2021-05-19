package main

import (
	"context"
	"time"

	"code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/services"
)

func main() {
	queue := bootstrap.NewOrderedQueue()

	feed := bootstrap.NewAxilisOfferFeed()

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	go queue.Start(context.Background())
	go feed.Start(ctx)

	feedProcessorService := services.NewFeedProcessorService(feed, queue)
	go feedProcessorService.Start(ctx)

	time.Sleep(time.Second*10)
}
