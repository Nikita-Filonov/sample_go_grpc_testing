package articlesservice

import (
	"context"
	articlesservice "github.com/Nikita-Filonov/sample_go_grpc_server/gen/proto"
	"github.com/dailymotion/allure-go"
	"github.com/onsi/gomega"
)

func (c *Client) GetArticle(
	g gomega.Gomega,
	request *articlesservice.GetArticleRequest,
) (response *articlesservice.GetArticleResponse) {
	allure.Step(allure.Description("Send ArticlesService.GetArticle request"), allure.Action(func() {
		response = c.getArticle(g, context.Background(), request)
	}))

	return response
}

func (c *Client) CreateArticle(
	g gomega.Gomega,
	request *articlesservice.CreateArticleRequest,
) (response *articlesservice.CreateArticleResponse) {
	allure.Step(allure.Description("Send ArticlesService.CreateArticle request"), allure.Action(func() {
		response = c.createArticle(g, context.Background(), request)
	}))

	return response
}

func (c *Client) UpdateArticle(
	g gomega.Gomega,
	request *articlesservice.UpdateArticleRequest,
) (response *articlesservice.UpdateArticleResponse) {
	allure.Step(allure.Description("Send ArticlesService.UpdateArticle request"), allure.Action(func() {
		response = c.updateArticle(g, context.Background(), request)
	}))

	return response
}

func (c *Client) DeleteArticle(
	g gomega.Gomega,
	request *articlesservice.DeleteArticleRequest,
) (response *articlesservice.DeleteArticleResponse) {
	allure.Step(allure.Description("Send ArticlesService.DeleteArticle request"), allure.Action(func() {
		response = c.deleteArticle(g, context.Background(), request)
	}))

	return response
}
