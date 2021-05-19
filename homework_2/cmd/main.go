package main

import (
	"fmt"

	"code-cadets-2021/homework_2/cmd/bootstrap"
	"code-cadets-2021/homework_2/internal/domain/services"
	"code-cadets-2021/homework_2/internal/tasks"
)

func main() {
	signalHandler := bootstrap.SignalHandler()

	feed1 := bootstrap.AxilisOfferFeed()
	feed2 := bootstrap.AxilisSecondOfferFeed()
	feeds := []services.Feed{feed1, feed2}

	queue := bootstrap.OrderedQueue()
	processingService := bootstrap.FeedProcessingService(feeds, queue)

	// blocking call, start "the application"
	tasks.RunTasks(signalHandler, feed1, feed2, queue, processingService)

	fmt.Println("program finished gracefully")
}
