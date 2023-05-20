package articlesservice

import (
	articlesservice "github.com/Nikita-Filonov/sample_go_grpc_server/gen/proto"
	"sample_go_grpc_testing/utils/config"
	"sample_go_grpc_testing/utils/logger"
	"sample_go_grpc_testing/utils/services/grpc"
)

type Client struct {
	articlesservice.ArticlesServiceClient
	logger *logger.CtxLogger
}

func NewClient(logger logger.Service, conf config.Config) (*Client, error) {
	conn, err := grpc.GetGrpcClient(conf.Articlesservice)

	if err != nil {
		return &Client{}, err
	}

	return &Client{
		logger:                logger.NewPrefix("ARTICLES_SERVICE.GRPC.CLIENT"),
		ArticlesServiceClient: articlesservice.NewArticlesServiceClient(conn),
	}, nil
}
