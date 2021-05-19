package bootstrap

import "code-cadets-2021/homework_2/internal/domain/services"

func FeedProcessingService(feed services.Feed, queue services.Queue) *services.FeedProcessorService {
	return services.NewFeedProcessorService(feed, queue)
}
