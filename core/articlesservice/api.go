package articlesservice

import (
	"context"
	articlesservice "github.com/Nikita-Filonov/sample_go_grpc_server/gen/proto"
	"github.com/onsi/gomega"
)

func (c *Client) createArticle(
	g gomega.Gomega,
	ctx context.Context,
	request *articlesservice.CreateArticleRequest,
) *articlesservice.CreateArticleResponse {
	c.logger.InfofJSON("CreateArticleRequest", request)

	res, err := c.ArticlesServiceClient.CreateArticle(ctx, request)
	g.Expect(err).ShouldNot(gomega.HaveOccurred(), "CreateArticle error")

	c.logger.InfofJSON("GetOperationsResponse", res)
	return res
}
