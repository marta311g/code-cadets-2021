package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/cmd/config"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/calcengine"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/calcengine/consumer"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/calcengine/handler"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/calcengine/publisher"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/mappers"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite"
)


func newEventUpdateConsumer(channel rabbitmq.Channel) *rabbitmq.EventUpdateConsumer {
	eventUpdateConsumer, err := rabbitmq.NewEventUpdateConsumer(
		channel,
		rabbitmq.ConsumerConfig{
			Queue:             config.Cfg.Rabbit.ConsumerEventUpdateQueue,
			DeclareDurable:    config.Cfg.Rabbit.DeclareDurable,
			DeclareAutoDelete: config.Cfg.Rabbit.DeclareAutoDelete,
			DeclareExclusive:  config.Cfg.Rabbit.DeclareExclusive,
			DeclareNoWait:     config.Cfg.Rabbit.DeclareNoWait,
			DeclareArgs:       nil,
			ConsumerName:      config.Cfg.Rabbit.ConsumerEventUpdateQueue,
			AutoAck:           config.Cfg.Rabbit.ConsumerAutoAck,
			Exclusive:         config.Cfg.Rabbit.ConsumerExclusive,
			NoLocal:           config.Cfg.Rabbit.ConsumerNoLocal,
			NoWait:            config.Cfg.Rabbit.ConsumerNoWait,
			Args:              nil,
		})
	if err != nil {
		panic(err)
	}
	return eventUpdateConsumer
}

func newBetConsumer(channel rabbitmq.Channel) *rabbitmq.BetConsumer {
	betConsumer, err := rabbitmq.NewBetConsumer(
		channel,
		rabbitmq.ConsumerConfig{
			Queue:             config.Cfg.Rabbit.ConsumerBetQueue,
			DeclareDurable:    config.Cfg.Rabbit.DeclareDurable,
			DeclareAutoDelete: config.Cfg.Rabbit.DeclareAutoDelete,
			DeclareExclusive:  config.Cfg.Rabbit.DeclareExclusive,
			DeclareNoWait:     config.Cfg.Rabbit.DeclareNoWait,
			DeclareArgs:       nil,
			ConsumerName:      config.Cfg.Rabbit.ConsumerBetQueue,
			AutoAck:           config.Cfg.Rabbit.ConsumerAutoAck,
			Exclusive:         config.Cfg.Rabbit.ConsumerExclusive,
			NoLocal:           config.Cfg.Rabbit.ConsumerNoLocal,
			NoWait:            config.Cfg.Rabbit.ConsumerNoWait,
			Args:              nil,
		})
	if err != nil {
		panic(err)
	}
	return betConsumer
}

func newConsumer(eventUpdateConsumer consumer.EventUpdateConsumer, betConsumer consumer.BetConsumer) *consumer.Consumer {
	return consumer.New(eventUpdateConsumer, betConsumer)
}

func newBetCalculatedMapper() *mappers.BetMapper {
	return mappers.NewBetCalculatedMapper()
}

func newCalcBetRepository(dbExecutor sqlite.DatabaseExecutor, betMapper sqlite.BetCalculatedMapper) *sqlite.CalcBetRepository {
	return sqlite.NewCalcBetRepository(dbExecutor, betMapper)
}

func newHandler(calcBetRepository handler.CalcBetRepository) *handler.Handler {
	return handler.New(calcBetRepository)
}

func newBetCalculatedPublisher(channel rabbitmq.Channel) *rabbitmq.BetCalculatedPublisher {
	betCalculatedPublisher, err := rabbitmq.NewBetCalculatedPublisher(
		channel,
		rabbitmq.PublisherConfig{
			Queue:             config.Cfg.Rabbit.PublisherBetCalculatedQueue,
			DeclareDurable:    config.Cfg.Rabbit.DeclareDurable,
			DeclareAutoDelete: config.Cfg.Rabbit.DeclareAutoDelete,
			DeclareExclusive:  config.Cfg.Rabbit.DeclareExclusive,
			DeclareNoWait:     config.Cfg.Rabbit.DeclareNoWait,
			DeclareArgs:       nil,
			Exchange:          config.Cfg.Rabbit.PublisherExchange,
			Mandatory:         config.Cfg.Rabbit.PublisherMandatory,
			Immediate:         config.Cfg.Rabbit.PublisherImmediate,
		},
	)
	if err != nil {
		panic(err)
	}
	return betCalculatedPublisher
}

func newPublisher(betCalculatedPublisher publisher.BetCalculatedPublisher) *publisher.Publisher {
	return publisher.New(betCalculatedPublisher)
}

func Engine(rabbitMqChannel rabbitmq.Channel, dbExecutor sqlite.DatabaseExecutor) *calcengine.Engine {
	eventUpdateConsumer := newEventUpdateConsumer(rabbitMqChannel)
	betConsumer := newBetConsumer(rabbitMqChannel)
	consumer := newConsumer(eventUpdateConsumer, betConsumer)

	betMapper := newBetCalculatedMapper()
	betRepository := newCalcBetRepository(dbExecutor, betMapper)
	handler := newHandler(betRepository)

	betPublisher := newBetCalculatedPublisher(rabbitMqChannel)
	publisher := newPublisher(betPublisher)

	return calcengine.New(consumer, handler, publisher)
}
