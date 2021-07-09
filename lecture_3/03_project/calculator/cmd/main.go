package main

import (
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/cmd/bootstrap"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/cmd/config"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/tasks"
	"log"
)

func main() {
	log.Println("Starting main")

	config.Load()

	rabbitMqChannel := bootstrap.RabbitMq()
	db := bootstrap.Sqlite()

	signalHandler := bootstrap.SignalHandler()
	engine := bootstrap.Engine(rabbitMqChannel, db)

	tasks.RunTasks(signalHandler, engine)
}
