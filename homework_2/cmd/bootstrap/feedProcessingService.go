package bootstrap

import "code-cadets-2021/homework_2/internal/domain/services"

func FeedProcessingService(feeds []services.Feed, queue services.Queue) *services.FeedProcessorService {
	return services.NewFeedProcessorService(feeds, queue)
}
