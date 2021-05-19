package bootstrap

import "code-cadets-2021/homework_2/internal/tasks"

func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
