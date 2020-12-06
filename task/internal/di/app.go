package di

import (
	"github.com/high-quality-sausages/sausage-task/task/internal/server"
	"github.com/high-quality-sausages/sausage-task/task/internal/service"
)

type App struct {
	svc *service.Service
}

func InitApp(svc *service.Service) (app *App) {
	server.Init(svc)
	app = &App{
		svc: svc,
	}

	return
}
