package container

import (
	"go.uber.org/dig"
	"sample_go_grpc_testing/core/articlesservice"
	"sample_go_grpc_testing/utils/config"
	"sample_go_grpc_testing/utils/logger"
)

func BuildContainer() (*Components, error) {
	c := dig.New()
	servicesConstructors := []interface{}{
		articlesservice.NewClient,
		config.NewConfig,
		logger.NewLoggerService,
	}

	for _, service := range servicesConstructors {
		err := c.Provide(service)
		if err != nil {
			return nil, err
		}
	}
	components, componentsError := initComponents(c)

	if componentsError != nil {
		return nil, componentsError
	}

	return components, nil
}
