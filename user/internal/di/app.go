package di

import (
	"github.com/high-quality-sausages/sausage-task/user/internal/service"
)

type App struct {
	svc *service.Service
}

func InitApp(svc *service.Service) (app *App) {
	app = &App{
		svc: svc,
	}
	return
}
