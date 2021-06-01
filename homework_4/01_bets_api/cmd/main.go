package main

import (
	"log"

	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/cmd/bootstrap"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/01_bets_api/internal/tasks"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	//rabbitMqChannel := bootstrap.RabbitMq()
	db := bootstrap.Sqlite()
	signalHandler := bootstrap.SignalHandler()
	api := bootstrap.Api(db)

	log.Println("Bootstrap finished. Event API is starting")

	tasks.RunTasks(signalHandler, api)

	log.Println("Event API finished gracefully")
}
