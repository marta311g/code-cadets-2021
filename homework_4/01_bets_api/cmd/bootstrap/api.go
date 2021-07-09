package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/api"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/domain/mappers"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/domain/services"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/infrastructure/sqlite"
)

func newBetMapper() *mappers.BetMapper {
	return mappers.NewBetMapper()
}

func newBetRepository(dbExecutor sqlite.DatabaseExecutor, mapper sqlite.BetMapper) *sqlite.BetRepository {
	return sqlite.NewBetRepository(dbExecutor, mapper)
}

func newBetService(repository sqlite.BetRepository) *services.BetService {
	return services.NewBetService(repository)
}

func newController(eventService controllers.DbService) *controllers.Controller {
	return controllers.NewController(eventService)
}

// Api bootstraps the http server.
func Api(dbExecutor sqlite.DatabaseExecutor) *api.WebServer {
	betMapper := newBetMapper()
	repository := newBetRepository(dbExecutor, betMapper)
	eventService := newBetService(*repository)
	controller := newController(eventService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
