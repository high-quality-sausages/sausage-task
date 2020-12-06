package di

import (
	"github.com/high-quality-sausages/sausage-task/user/internal/dao"
	"github.com/high-quality-sausages/sausage-task/user/internal/service"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(dao.NewDB)
	container.Provide(dao.New)
	container.Provide(service.New)
	container.Provide(InitApp)

	return container
}
