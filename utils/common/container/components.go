package container

import (
	"go.uber.org/dig"
	"sample_go_grpc_testing/core/articlesservice"
	"sample_go_grpc_testing/utils/config"
	"sample_go_grpc_testing/utils/logger"
)

type Components struct {
	ArticlesService *articlesservice.Client

	Logger logger.Service
	Config config.Config
}

func initComponents(c *dig.Container) (*Components, error) {
	var err error
	components := Components{}

	err = c.Invoke(func(
		articlesService *articlesservice.Client,

		logger logger.Service,
		conf config.Config,
	) {
		components.Config = conf
		components.Logger = logger

		components.ArticlesService = articlesService
	})

	if err != nil {
		return nil, err
	}

	return &components, nil
}
