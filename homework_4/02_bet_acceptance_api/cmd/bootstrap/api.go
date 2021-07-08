package bootstrap

import (
	"github.com/streadway/amqp"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_acceptance_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_acceptance_api/internal/api"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_acceptance_api/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_acceptance_api/internal/api/controllers/validators"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_acceptance_api/internal/domain/services"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bet_acceptance_api/internal/infrastructure/rabbitmq"
)

func newBetValidator() *validators.BetValidator {
	return validators.NewBetValidator()
}

func newBetPublisher(publisher rabbitmq.QueuePublisher) *rabbitmq.BetReceivedPublisher {
	return rabbitmq.NewBetReceivedPublisher(
		config.Cfg.Rabbit.PublisherExchange,
		config.Cfg.Rabbit.PublisherBetReceivedQueueQueue,
		config.Cfg.Rabbit.PublisherMandatory,
		config.Cfg.Rabbit.PublisherImmediate,
		publisher,
	)
}

func newBetService(publisher services.BetReceivedPublisher) *services.BetService {
	return services.NewBetService(publisher)
}

func newController(betValidator controllers.BetValidator, eventService controllers.BetService) *controllers.Controller {
	return controllers.NewController(betValidator, eventService)
}

// Api bootstraps the http server.
func Api(rabbitMqChannel *amqp.Channel) *api.WebServer {
	betValidator := newBetValidator()
	betPublisher := newBetPublisher(rabbitMqChannel)
	betService := newBetService(betPublisher)
	controller := newController(betValidator, betService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
