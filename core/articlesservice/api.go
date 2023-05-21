package articlesservice

import (
	"context"
	articlesservice "github.com/Nikita-Filonov/sample_go_grpc_server/gen/proto"
	"github.com/onsi/gomega"
)

func (c *Client) getArticle(
	g gomega.Gomega,
	ctx context.Context,
	request *articlesservice.GetArticleRequest,
) *articlesservice.GetArticleResponse {
	c.logger.InfofJSON("GetArticleRequest", request)

	res, err := c.ArticlesServiceClient.GetArticle(ctx, request)
	g.Expect(err).ShouldNot(gomega.HaveOccurred(), "GetArticle error")

	c.logger.InfofJSON("GetArticleResponse", res)
	return res
}

func (c *Client) createArticle(
	g gomega.Gomega,
	ctx context.Context,
	request *articlesservice.CreateArticleRequest,
) *articlesservice.CreateArticleResponse {
	c.logger.InfofJSON("CreateArticleRequest", request)

	res, err := c.ArticlesServiceClient.CreateArticle(ctx, request)
	g.Expect(err).ShouldNot(gomega.HaveOccurred(), "CreateArticle error")

	c.logger.InfofJSON("CreateArticleResponse", res)
	return res
}

func (c *Client) updateArticle(
	g gomega.Gomega,
	ctx context.Context,
	request *articlesservice.UpdateArticleRequest,
) *articlesservice.UpdateArticleResponse {
	c.logger.InfofJSON("UpdateArticleRequest", request)

	res, err := c.ArticlesServiceClient.UpdateArticle(ctx, request)
	g.Expect(err).ShouldNot(gomega.HaveOccurred(), "UpdateArticle error")

	c.logger.InfofJSON("UpdateArticleResponse", res)
	return res
}

func (c *Client) deleteArticle(
	g gomega.Gomega,
	ctx context.Context,
	request *articlesservice.DeleteArticleRequest,
) *articlesservice.DeleteArticleResponse {
	c.logger.InfofJSON("DeleteArticleRequest", request)

	res, err := c.ArticlesServiceClient.DeleteArticle(ctx, request)
	g.Expect(err).ShouldNot(gomega.HaveOccurred(), "DeleteArticle error")

	c.logger.InfofJSON("DeleteArticleResponse", res)
	return res
}
