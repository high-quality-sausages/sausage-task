package di

import (
	"go.uber.org/dig"

	"github.com/high-quality-sausages/sausage-task/task/internal/dao"
	"github.com/high-quality-sausages/sausage-task/task/internal/service"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(dao.NewDB)
	container.Provide(dao.New)
	container.Provide(service.New)
	container.Provide(InitApp)

	return container
}
