package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/api"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/domain/services"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/infrastructure/sqlite"
)

func newBetService(dbExecutor sqlite.DatabaseExecutor) *services.BetService {
	return services.NewBetService(dbExecutor)
}

func newController(eventService controllers.DbService, dbExecutor sqlite.DatabaseExecutor) *controllers.Controller {
	return controllers.NewController(eventService, dbExecutor)
}

// Api bootstraps the http server.
func Api(dbExecutor sqlite.DatabaseExecutor) *api.WebServer {
	eventService := newBetService(dbExecutor)
	controller := newController(eventService, dbExecutor)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
